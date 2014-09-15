// Copyright © 2010-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
    "fmt"
    "log"
    "os"
    "flag"
    "strings"
)

var help, bar bool

func init() {
    // set up command line parsing
    const (
        help_usage = "show help"
        bar_usage  = "print bar above/below number"
    )
    flag.BoolVar(&help, "help", false, help_usage)
    flag.BoolVar(&help, "h", false, help_usage+" (shorthand)")
    flag.BoolVar(&bar, "bar", false, bar_usage)
    flag.BoolVar(&bar, "b", false, bar_usage+" (shorthand)")
}

func main() {
    // show usage and exit if
    //   -h or --help was passed
    //   the flags failed to parse
    //   there isn't a single arg left
    flag.Parse()
    if help || !flag.Parsed() || flag.NArg() != 1 {
        usage()
        os.Exit(1)
    }

    stringOfDigits := flag.Arg(0)
    barlength := getbarlength(stringOfDigits)
    if bar {
        fmt.Println(strings.Repeat("*", barlength))
    }
    for row := range bigDigits[0] {
        line := ""
        for column := range stringOfDigits {
            digit := stringOfDigits[column] - '0'
            if 0 <= digit && digit <= 9 {
                line += bigDigits[digit][row] + "  "
            } else {
                log.Fatal("invalid whole number")
            }
        }
        fmt.Println(line)
    }
    if bar {
        fmt.Println(strings.Repeat("*", barlength))
    }
}

func usage() {
    fmt.Println("usage: bigdigits [-b|--bar] whole-number")
    fmt.Println("-b --bar draw an underbar and an overbar")
}

func getbarlength(stringOfDigits string) int {
    var length int
    for column := range stringOfDigits {
        var maxwidth int
        digit := stringOfDigits[column] - '0'
        if 0 <= digit && digit <= 9 {
            for row := range bigDigits[digit] {
                if len(bigDigits[digit][row]) > maxwidth {
                    maxwidth = len(bigDigits[digit][row])
                }
            }
        } else {
            log.Fatal("invalid whole number")
        }
        length += (maxwidth + 2)
    }
    return (length - 2)
}

var bigDigits = [][]string{
    {"  000  ",
     " 0   0 ",
     "0     0",
     "0     0",
     "0     0",
     " 0   0 ",
     "  000  "},
    {" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
    {" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
    {" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
    {"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
        "   4  "},
    {"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
    {" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
    {"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
    {" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
    {" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}
