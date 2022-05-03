/*
Copyright © 2022 Augustyn Chmiel <chmielua@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

The cobra package also supports subcommands, which are commands associated with specific commands
like go run main.go command list. Try to implement a utility with subcommands.

(Page 448).
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/nXnUs25/golinuxcmds/wcgo"
	"github.com/spf13/cobra"
)

// wcCmd represents the wc command
var wcCmd = &cobra.Command{
	Use:   "wc [file ...]",
	Short: "wc - word, line, character, and byte count",
	Long: `The wc utility displays the number of lines, words, and bytes contained in each input file, 
or standard input (if no file is specified) to the standard output.  A line
is defined as a string of characters delimited by a ⟨newline⟩ character.  
Characters beyond the final ⟨newline⟩ character will not be included in the line count.

A word is defined as a string of characters delimited by white space characters.  
White space characters are the set of characters for which the iswspace(3) function
returns true.  If more than one input file is specified, a line of cumulative counts for 
all the files is displayed on a separate line after the output for the last file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		lines, err := cmd.Flags().GetBool("lines")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		words, err := cmd.Flags().GetBool("words")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		chars, err := cmd.Flags().GetBool("chars")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		bytes, err := cmd.Flags().GetBool("bytes")
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if len(args) > 0 {
			for _, x := range args {
				wc := &wcgo.Wc{}
				wcgo.Filename = x
				wc.ReadFile()
				switch {
				case lines:
					fmt.Fprintln(os.Stdout, wc.GetLinesCount())
				case chars:
					fmt.Fprintln(os.Stdout, wc.GetCharsCount())
				case words:
					fmt.Fprintln(os.Stdout, wc.GetWordsCount())
				case bytes:
					fmt.Fprintln(os.Stdout, wc.GetBytesCount())
				default:
					fmt.Fprintln(os.Stdout, wc)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(wcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	wcCmd.Flags().BoolP("lines", "l", false, "The number of lines in each input file is written to the standard output.")
	wcCmd.Flags().BoolP("words", "w", false, "The number of words in each input file is written to the standard output.")
	wcCmd.Flags().BoolP("chars", "c", false, "The number of characters in each input file is written to the standard output.")
	wcCmd.Flags().BoolP("bytes", "b", false, "The number of bytes in each input file is written to the standard output.")
}
