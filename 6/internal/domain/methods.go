package domain

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func (d *Domain) Grep(flags Flags, pattern string, files []*os.File) error {
	if flags.Fixed {
		for _, file := range files {
			err := d.GrepFixed(flags, pattern, file)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *Domain) GrepFixed(flags Flags, pattern string, file *os.File) error {
	buf := bytes.Buffer{}
	scanner := bufio.NewScanner(file)
	var lineCount int

	for scanner.Scan() {
		lineCount++
		line := scanner.Text()

		if flags.IgnoreCase {
			line = strings.ToLower(line)
		}

		if line == pattern {
			colored := color.GreenString(line)
			if flags.LineNum {
				buf.Write([]byte(fmt.Sprintf("%d:%s\n", lineCount, colored)))
			} else {
				buf.Write([]byte(fmt.Sprintf("%s\n", colored)))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		d.log.Error(err.Error())
		return err
	}

	fmt.Println(strings.TrimSuffix(buf.String(), "\n"))
	return nil
}

// func (d *Domain) GrepRegexp(flags Flags, pattern string, file *os.File) error {
// 	buf := bytes.Buffer{}
// 	scanner := bufio.NewScanner(file)
// 	var lineCount int
// }
