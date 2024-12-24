package main

import (
	"log"
	"log/slog"
	"wildberries-l2/grep/internal/view"

	"github.com/spf13/cobra"
)

func main() {

	cmd := &cobra.Command{
		Use:   "grep [flags] pattern [files...]",
		Short: "Search for a pattern in files",
		Run:   runGrep,
	}

	initCmd(cmd)

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

}

func runGrep(cmd *cobra.Command, args []string) {
	logger := slog.New(slog.NewTextHandler(log.Writer(), nil))
	view := view.NewView(*logger)
	view.Grep(cmd, args)
}

func initCmd(cmd *cobra.Command) {
	cmd.Flags().IntP("after", "A", 0, "Print N lines after match")
	cmd.Flags().IntP("before", "B", 0, "Print N lines before match")
	cmd.Flags().IntP("context", "C", 0, "Print N lines before and after match")
	cmd.Flags().IntP("count", "c", 0, "Print the count of matching lines")
	cmd.Flags().BoolP("ignore-case", "i", false, "Ignore case distinctions")
	cmd.Flags().BoolP("invert", "v", false, "Invert match")
	cmd.Flags().BoolP("fixed", "F", false, "Match not a pattern but a fixed string")
	cmd.Flags().BoolP("line num", "n", false, "Print line number")
}
