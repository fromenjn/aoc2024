package utils

import (
	"fmt"
	"strconv"
)

type DiskUnit struct {
	FileID int
}

func StringToDiskUnits(l string) []DiskUnit {
	units := make([]DiskUnit, 0)
	FileID := 0
	for i := 0; i < len(l); i++ {
		lenBlock, err := strconv.Atoi(string(l[i]))
		if err != nil {
			fmt.Printf("Skipping invalid char in line: %c\n", l[i])
		}
		if i%2 == 0 {
			for j := 0; j < lenBlock; j++ {
				units = append(units, DiskUnit{FileID: FileID})
			}
			FileID++
		} else {
			for j := 0; j < lenBlock; j++ {
				units = append(units, DiskUnit{FileID: -1})
			}
		}
	}
	return units
}

func FragmentDisk(units []DiskUnit) {
	idxAvail := 0
	for i := len(units) - 1; i >= 0; i-- {
		fileID := units[i].FileID
		if fileID >= 0 && idxAvail < i {
			//fmt.Print(fileID)
			for (idxAvail < len(units)) && (units[idxAvail].FileID >= 0) {
				idxAvail++
			}
			if (idxAvail < len(units)) && (units[idxAvail].FileID == -1) && idxAvail < i {
				fmt.Printf("%d is now %d\n", idxAvail, units[i].FileID)
				units[idxAvail].FileID = units[i].FileID
				units[i].FileID = -1
			}
		}
	}
	fmt.Println()
	PrintDiskUnits(units)
}

func GetUsedBlockSize(units []DiskUnit, idxEnd int) int {
	blockSize := 0
	fileID := units[idxEnd].FileID
	for i := idxEnd; i >= 0 && units[i].FileID == fileID; i-- {
		blockSize++
	}
	return blockSize
}

func GetAvailBlockSize(units []DiskUnit, idxStart int) int {
	blockSize := 0
	for i := idxStart; i < len(units) && units[i].FileID == -1; i++ {
		blockSize++
	}
	return blockSize
}

func SwapBlockFile(units []DiskUnit, idxEndUsedBlock int, idxStartAvailBlock int, usedBlockSize int) {
	fileID := units[idxEndUsedBlock].FileID
	for i := 0; i < usedBlockSize; i++ {
		units[idxEndUsedBlock-i].FileID = -1
		units[idxStartAvailBlock+i].FileID = fileID
	}
}

func DefragmentDisk(units []DiskUnit) {
	//fmt.Printf("lenUnits: %d\n", len(units))
	i := len(units) - 1
	for i >= 0 {
		currentFileID := units[i].FileID
		if currentFileID >= 0 {
			usedBlockSize := GetUsedBlockSize(units, i)
			//fmt.Printf("Used Block Size for id %d: %d\n", units[i].FileID, usedBlockSize)
			for j := 0; j < i; j++ {
				availBlockSize := GetAvailBlockSize(units, j)
				if availBlockSize >= usedBlockSize {
					//fmt.Printf("Swapping %d with %d for len %d\n", i, j, usedBlockSize)
					SwapBlockFile(units, i, j, usedBlockSize)
					break
				}
			}
			for i >= 0 && currentFileID == units[i].FileID {
				i--
			}
		} else {
			i--
		}

	}
}

func PerformChecksum(units []DiskUnit) int {
	checkSum := 0
	for i := 0; i < len(units); i++ {
		if units[i].FileID != -1 {
			checkSum += i * units[i].FileID
		}
	}
	return checkSum
}

func PrintDiskUnits(units []DiskUnit) {
	for i := 0; i < len(units); i++ {
		if units[i].FileID == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(units[i].FileID)
		}
	}
	fmt.Println()
}
