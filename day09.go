package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Item struct {
	id     int
	length int
}
type Items []Item

func (list Items) PopFront() (Items, Item) {
	return list[1:], list[0]
}
func (list Items) PopBack() (Items, Item) {
	return list[:len(list)-1], list[len(list)-1]
}
func (list Items) Move(from, to int) Items {
	var temp Items
	mv := list[from]
	list = slices.Delete(list, from, from+1)
	temp = append(temp, list[0:to]...)
	temp = append(temp, mv)
	temp = append(temp, list[to+1:]...)
	return list
}
func (list Items) FindId(id int) int {
	return slices.IndexFunc(list, func(i Item) bool {
		return i.id == id
	})
}

func day09(part int, file *os.File) {

	var files Items
	var freeSpace Items

	tick := 0
	id := 0
	r := bufio.NewReader(file)
	for {
		b, err := r.ReadByte()
		if err != nil {
			break
		}
		if b == '\n' {
			continue
		}
		num := int(b) - 48
		if tick == 0 {
			files = append(files, Item{id, num})
		} else {
			freeSpace = append(freeSpace, Item{id, num})
		}
		if tick == 1 {
			id++
		}
		tick = (tick + 1) % 2
	}
	/*
		for i := 0; i < len(files); i++ {
			fmt.Printf("%s", strings.Repeat(strconv.Itoa(i), files[i].length))
			if i == len(freeSpace) {
				continue
			}
			fmt.Printf("%s", strings.Repeat(".", freeSpace[i].length))
		}
		fmt.Printf("\n")
	*/

	var disk Items

	if part == 1 {
		var newFile, fspace, sfile Item
		for {
			if len(files) == 0 && sfile.length == 0 {
				break
			}
			if fspace.length == 0 && len(files) > 0 {
				files, newFile = files.PopFront()
				disk = append(disk, newFile)
				if len(freeSpace) > 0 {
					freeSpace, fspace = freeSpace.PopFront()
				}
				continue
			}
			if len(files) == 0 && sfile.length > 0 {
				disk = append(disk, sfile)
				break
			}
			if sfile.length == 0 && len(files) > 0 {
				files, sfile = files.PopBack()
			}

			diff := fspace.length - sfile.length
			if diff < 0 {
				disk = append(disk, Item{sfile.id, fspace.length})
				sfile.length = sfile.length - fspace.length
				fspace.length = 0
			} else {
				disk = append(disk, sfile)
				fspace.length -= sfile.length
				sfile.length = 0
			}
		}
	}

	if part == 2 {
		max := 0
		for i := 0; i < len(files); i++ {
			disk = append(disk, files[i])
			if files[i].id > max {
				max = files[i].id
			}
			if i >= len(freeSpace) {
				continue
			}
			disk = append(disk, Item{-1, freeSpace[i].length})
		}

		for curr := max; curr > 0; curr-- {
			id := disk.FindId(curr)
			if id == -1 {
				continue
			}
			f := disk[id]
			for idx, d := range disk {
				if idx >= id {
					break
				}
				if d.id != -1 {
					continue
				}
				if d.length < f.length {
					continue
				}
				diff := d.length - f.length
				disk[idx].length = diff
				disk[id].id = -1
				newSeq := Items{{-1, 0}, f}
				disk = slices.Insert(disk, idx, newSeq...)
				break
			}
		}
	}

	id = 0
	checksum := 0
	for _, f := range disk {
		if f.id == -1 {
			id += f.length
			continue
		}
		for n := 0; n < f.length; n++ {
			checksum += id * f.id
			id++
		}
	}

	fmt.Printf("Day 09 part %d: %d\n", part, checksum)

}
