package view

import (
	"os"
	"wildberries-l2/grep/internal/domain"

	"github.com/spf13/cobra"
)

func (v *View) Grep(cmd *cobra.Command, args []string) {

	if len(args) < 2 {
		v.log.Error("Not enough arguments")
		v.log.Error("usage: grep [OPTIONS] [PATTERN] [FILE]")
		return
	}

	flags := domain.Flags{}

	// throws non name error on the left side
	// of := without this line
	var err error

	flags.After, err = cmd.Flags().GetInt("after")
	if err != nil {
		v.log.Error(err.Error())
		return
	}
	flags.Before, err = cmd.Flags().GetInt("before")
	if err != nil {
		v.log.Error(err.Error())
		return
	}
	flags.Context, err = cmd.Flags().GetInt("context")
	if err != nil {
		v.log.Error(err.Error())
		return
	}

	flags.Count, err = cmd.Flags().GetInt("count")
	if err != nil {
		v.log.Error(err.Error())
		return
	}

	flags.IgnoreCase, err = cmd.Flags().GetBool("ignore-case")
	if err != nil {
		v.log.Error(err.Error())
		return
	}

	flags.Invert, err = cmd.Flags().GetBool("invert")
	if err != nil {
		v.log.Error(err.Error())
		return
	}

	flags.Fixed, err = cmd.Flags().GetBool("fixed")
	if err != nil {
		v.log.Error(err.Error())
		return
	}

	flags.LineNum, err = cmd.Flags().GetBool("line num")
	if err != nil {
		v.log.Error(err.Error())
		return
	}

	var files []*os.File

	pattern := args[0]
	filenames := args[1:]

	for _, filename := range filenames {
		file, err := os.Open(filename)

		if err != nil {
			v.log.Error(err.Error())
			return
		}

		defer file.Close()

		files = append(files, file)
	}

	err = v.domain.Grep(flags, pattern, files)

	if err != nil {
		v.log.Error(err.Error())
		return
	}

}
