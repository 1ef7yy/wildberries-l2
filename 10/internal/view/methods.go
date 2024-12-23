package view

import "fmt"

func (v *View) Wget(args []string) {
	if len(args) == 0 {
		fmt.Println("usage: wget [URL] [DESTINATION]")
		return
	}

	var dest string

	url := args[1]

	if len(args) == 2 {
		dest = ""
	} else {
		dest = args[2]
	}

	val, err := v.domain.Wget(url, dest)
	if err != nil {
		v.log.Error(err.Error())
		return
	}

	if dest == "" {
		fmt.Println(val)
		return
	}

	fmt.Println("File saved to", dest)
	return
}
