package writer

import (
	"flag"
	"fmt"
	"os"
)

const Usage = `Usage: writefile -size SIZE_BYTES PATH
Example: writefile -size 1000 zeroes.dat`

func WriteToFile(path string, data []byte) error {
	e := os.WriteFile(path, data, 0o600)
	if e != nil {
		return e
	}
	return os.Chmod(path, 0o600)
}

// func WriteZeroesToFile(numberOfZeroes int, path string) error {
// 	data := []byte{}
// 	for i := range numberOfZeroes {
// 		data = append(data, byte(i))
// 	}
// 	return WriteToFile(path, data)
// }

func Main() int {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		return 0
	}
	size := flag.Int("size", 0, "Size in bytes")
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, Usage)
		return 1
	}
	err := WriteToFile(flag.Args()[0], make([]byte, *size))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	return 0
}
