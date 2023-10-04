package get

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

// GetInfo retrieves information about a file or directory.
// It categorizes the file based on flags and handles symbolic links.
func GetInfo(file File, flagsToUse Flags, index int) (File, error) {
	// Open the current working directory.
	dir, err := os.Open(file.CWD)
	if err != nil {
		return file, err // NotThere
	}
	defer dir.Close()

	fileInfo, _ := os.Lstat(file.CWD)
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		// Check if it's a symbolic link --Cyan
		file.NotFolder = true // not a directory or folder
	}

	var entries []string
	if !file.NotFolder {
		entries, err = dir.Readdirnames(0) // Read all entries in the directory.
		if err != nil {
			file.NotFolder = true // not a directory or folder
		}
	}

	if !file.NotFolder {
		// Create a slice to store the names of both files and directories.
		var names []string
		var hiddenNames []string

		// Collect both file and directory names (excluding hidden files).
		for _, name := range entries {
			if !strings.HasPrefix(name, ".") { // Exclude hidden files.
				names = append(names, name)
			} else if strings.HasPrefix(name, ".") {
				hiddenNames = append(hiddenNames, name)
			} else {
				continue
			}
		}
		if flagsToUse.Flag_a {
			hiddenNames = append(hiddenNames, ".")
			hiddenNames = append(hiddenNames, "..")
			names = append(names, hiddenNames...)
		}

		// Sort the names as if they were lowercase but display the actual values.
		sort.SliceStable(names, func(i, j int) bool {
			if flagsToUse.Flag_t {
				filePath1 := filepath.Join(file.CWD, names[i])
				filePath2 := filepath.Join(file.CWD, names[j])

				fileInfo1, err1 := os.Lstat(filePath1)
				fileInfo2, err2 := os.Lstat(filePath2)

				// Handle errors, e.g., if there's a problem getting file info.
				if err1 != nil || err2 != nil {
					return false // You may want to handle errors differently.
				}
				// Compare files by ModTime.
				if flagsToUse.Flag_r {
					// If Flag_r is true, reverse the order.
					return fileInfo1.ModTime().Before(fileInfo2.ModTime())
				}
				return fileInfo1.ModTime().After(fileInfo2.ModTime())

			}

			if flagsToUse.Flag_r {
				// If Flag_r is true, reverse the order of names.
				return strings.ToLower(names[i]) > strings.ToLower(names[j])
			}

			return strings.ToLower(names[i]) < strings.ToLower(names[j])
		})

		file.Names, file.HiddenNames = names, hiddenNames
	}

	file = AppendData(file)
	return file, nil
}

// GetFileType categorizes a file based on its type and assigns a color.
// It also handles symbolic links and updates arguments if the -R flag is used.
func GetFileType(file File, args []string, flagsToUse Flags, mainRoot string, index int) (File, []string, error) {
	var fileTypes []string
	var fileColor []string
	var softLinks []string

	if file.FileName != "." && file.FileName != ".." {
		file.FileName = TrimFliePath(file.CWD, mainRoot)
	}

	// if file.NotFolder
	for _, name := range file.Names {
		// Join the subfolder path and file name to get the full file path.
		filePath := filepath.Join(file.CWD, name)

		// Get file information
		fileInfo, err := os.Lstat(filePath)
		// fileInfo, err := os.Stat(filePath) /// does not recognize symbolic links "softlinks""
		if err != nil {
			// fmt.Println("iiiiii", err) ////////////////////////////////////////////////////////////////////////
			return file, args, err
		}

		// List of common graphic image file extensions.
		imageExtensions := ".jpg .jpeg .png .gif .bmp .webp .svg .tiff .ico"
		// List of common archive file extensions.
		archiveExtensions := ".zip .tar .gz .bz2 .xz .rar .7z"

		// Check if it's a directory or folder --Blue
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			// Check if it's a symbolic link --Cyan
			fileTypes = append(fileTypes, "Link")
			fileColor = append(fileColor, MyColor("Link"))

			link, _ := os.Readlink(filePath)
			softLinks = append(softLinks, link)

		} else if fileInfo.IsDir() {
			fileTypes = append(fileTypes, "Directory")
			fileColor = append(fileColor, MyColor("Directory"))
			softLinks = append(softLinks, "")

			if flagsToUse.Flag_R {
				if name == "." || name == ".." {
					continue
				}
				args = append(args, filePath)
			}
		} else if fileInfo.Mode()&0111 != 0 {
			// Check if it's an executable file --Green
			fileTypes = append(fileTypes, "Executable")
			fileColor = append(fileColor, MyColor("Executable"))
			softLinks = append(softLinks, "")
		} else if strings.Contains(imageExtensions, strings.ToLower(FindExt(filePath))) {
			// Check if it's an Image --Pink
			fileTypes = append(fileTypes, "Image")
			fileColor = append(fileColor, MyColor("Image"))
			softLinks = append(softLinks, "")
		} else if strings.Contains(archiveExtensions, strings.ToLower(FindExt(filePath))) {
			// Check if it's an archive file --Red
			fileTypes = append(fileTypes, "Archive")
			fileColor = append(fileColor, MyColor("Archive"))
			softLinks = append(softLinks, "")
		} else {
			// Default: Unrecognized type
			fileTypes = append(fileTypes, "Unrecognized")
			fileColor = append(fileColor, MyColor("Unrecognized"))
			softLinks = append(softLinks, "")
		}
	}
	file.Types = fileTypes
	file.FileColor = fileColor
	file.SoftLinks = softLinks

	return file, args, nil
}

// AppendData collects data from all files to save it into the File struct.
func AppendData(file File) File {
	var ModiTime []string
	var Permission []string
	var Owner []string
	var Group []string
	var SizeKB []int
	var TotalSizeKB int
	var Size []int64
	var Hardlinks []int
	var Month []string
	var Day []int

	if file.NotFolder {
		fileInfo, _ := os.Lstat(file.CWD)

		if fileInfo.Mode()&os.ModeSymlink != 0 {
			// Check if it's a symbolic link --Cyan
			file.Types = append(file.Types, "Link")
			file.FileColor = append(file.FileColor, MyColor("Link"))

			link, _ := os.Readlink(file.CWD)
			file.SoftLinks = append(file.SoftLinks, link)
		} else {
			file.SoftLinks = append(file.SoftLinks, "")
		}

		timeToAppend := fmt.Sprintf("%+03d:%+03d", fileInfo.ModTime().Hour(), fileInfo.ModTime().Minute())
		timeToAppend = strings.Replace(timeToAppend, "+", "", -1)
		ModiTime = append(ModiTime, timeToAppend)
		Permission = append(Permission, fmt.Sprintf("%v", fileInfo.Mode()))

		Day = append(Day, fileInfo.ModTime().Day())
		Month = append(Month, fileInfo.ModTime().Format("Jan "))

		Size = append(Size, fileInfo.Size())

		if stat, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
			UID, _ := user.LookupId(strconv.Itoa(int(stat.Uid)))
			GID, _ := user.LookupGroupId(strconv.Itoa(int(stat.Gid)))
			SizeKB = append(SizeKB, int(stat.Blocks/2))
			TotalSizeKB = TotalSizeKB + int(stat.Blocks/2)
			Owner = append(Owner, UID.Username)
			Group = append(Group, GID.Name)
			Hardlinks = append(Hardlinks, int(stat.Nlink))
		}

	} else {
		for _, name := range file.Names {
			filePath := filepath.Join(file.CWD, name)
			fileInfo, _ := os.Lstat(filePath)

			timeToAppend := fmt.Sprintf("%+03d:%+03d", fileInfo.ModTime().Hour(), fileInfo.ModTime().Minute())
			timeToAppend = strings.Replace(timeToAppend, "+", "", -1)
			ModiTime = append(ModiTime, timeToAppend)
			Permission = append(Permission, fmt.Sprintf("%v", fileInfo.Mode()))

			Day = append(Day, fileInfo.ModTime().Day())
			Month = append(Month, fileInfo.ModTime().Format("Jan "))
			Size = append(Size, fileInfo.Size())
			file.SoftLinks = append(file.SoftLinks, "")
			if stat, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
				UID, _ := user.LookupId(strconv.Itoa(int(stat.Uid)))
				GID, _ := user.LookupGroupId(strconv.Itoa(int(stat.Gid)))
				SizeKB = append(SizeKB, int(stat.Blocks/2))
				TotalSizeKB = TotalSizeKB + int(stat.Blocks/2)
				Owner = append(Owner, UID.Username)
				Group = append(Group, GID.Name)
				Hardlinks = append(Hardlinks, int(stat.Nlink))
			}
		}
	}

	file.ModificationTime = ModiTime
	file.Permission = Permission
	file.Owner = Owner
	file.Group = Group
	file.SizeKB = SizeKB
	file.Size = Size
	file.TotalSize = TotalSizeKB
	file.Hardlinks = Hardlinks
	file.Day = Day
	file.Month = Month

	return file
}
