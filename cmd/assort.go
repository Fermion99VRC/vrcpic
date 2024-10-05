package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var pictureDirectory string
var targetDirectory string
var keepPicture bool

// assortCmd represents the assort command
var assortCmd = &cobra.Command{
	Use:   "assort",
	Short: "assort screenshots of VRChat",
	Long: `Move screenshots of VRChat to <to>/YYYY-MM/MMDD.
<to> is given with the flag "--to".
Note that the screenshots taken after 12:00 are moved to the next-day directory.
For example, the screenshot taken at 18:00 on 2024-01-01 will be moved to <to>/2024-01/0102.`,
	Run: func(cmd *cobra.Command, args []string) {
		pattern := filepath.Join(pictureDirectory, "*", "*.png")
		paths, err := filepath.Glob(pattern)
		if err != nil {
			// TODO: Find the best practice of error handling
			fmt.Printf("Failed to glob files of pictures with %q: %v\n", pattern, err)
			os.Exit(1)
		}
		for _, path := range paths {
			filename := filepath.Base(path)
			datetime, err := ParseFilename(filename)
			if err != nil {
				fmt.Printf("Failed to parse %q: %v\n", filename, err)
				continue
			}
			dirname := GetDirectoryToMove(datetime)
			newpath := filepath.Join(targetDirectory, dirname, filename)
			os.MkdirAll(filepath.Dir(newpath), 0664)
			if err := CopyFile(path, newpath); err != nil {
				fmt.Printf("Failed to copy file: %v\n", err)
				continue
			}
			if !keepPicture {
				if err := os.Remove(path); err != nil {
					fmt.Printf("Couldn't remove source file: %v", err)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(assortCmd)
	assortCmd.Flags().StringVarP(&pictureDirectory, "from", "f", ".", "directory where VRChat app saves screenshots")
	assortCmd.Flags().StringVarP(&targetDirectory, "to", "t", ".", "directory to which screenshots are moved")
	assortCmd.Flags().BoolVar(&keepPicture, "keep", false, "keep screenshots in the original directory")
}

// Parse the filename of screenshot to return its date and time.
func ParseFilename(filename string) (time.Time, error) {
	splitted := strings.Split(filename, "_")
	if len(splitted) != 4 {
		/*
			The spec of filename is "VRChat_YYYY-MM-DD_hh-mm-ss.sss_WIDTHxHEIGHT.png"
			If this function returns error, VRChat may change the specifications and you will have to
			fix this function.
		*/
		return time.Time{}, fmt.Errorf("The filename doesn't match the spec: %q", filename)
	}
	datestr := splitted[1] + "T" + strings.Replace(splitted[2], "-", ":", -1) + "Z"
	return time.Parse(time.RFC3339Nano, datestr)
}

// Return the directory to move according to the time when the picture is taken.
func GetDirectoryToMove(datetime time.Time) string {
	if datetime.Hour() >= 12 {
		datetime = datetime.AddDate(0, 0, 1)
	}
	root := fmt.Sprintf("%d-%02d", datetime.Year(), datetime.Month())
	dirname := fmt.Sprintf("%02d%02d", datetime.Month(), datetime.Day())
	return filepath.Join(root, dirname)
}

// Copy a source file to the destination.
func CopyFile(src, dest string) error {
	inputf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %v", err)
	}
	defer inputf.Close()

	outputf, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("Couldn't create dest file: %v", err)
	}
	defer outputf.Close()

	_, err = io.Copy(outputf, inputf)
	if err != nil {
		return fmt.Errorf("Couldn't copy from src to dest: %v", err)
	}
	return nil
}
