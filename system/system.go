package system

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

const (
	ColorBlack  = "\u001b[30m"
	ColorRed    = "\u001b[31m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorBlue   = "\u001b[34m"
	ColorReset  = "\u001b[0m"
)

func Warning(message string) {
	fmt.Println(string(ColorYellow), message, string(ColorReset))
}

func Break(message string) {
	fmt.Println(string(ColorRed), message, string(ColorReset))
	os.Exit(0)
}

func Pass(message string) {
	fmt.Println(string(ColorGreen), message, string(ColorReset))

}

func Help(message string) {
	fmt.Println(string(ColorBlue), message, string(ColorReset))
}

func Red(message string) {
	fmt.Println(string(ColorRed), message, string(ColorReset))
}

func Count_pos() int {
	var arg_len int = len(os.Args[1:])
	return arg_len
}

// help := "-h"
// version := "-v"
// initialize := "-i"
// test -t
// read := "-r"
// write := "-w"
// destroy := "-d"
// uninstall := "--uninstall"
// update := "--update"

func Encrypt_san(arguments int, filename string, object_name string, object_owner string) string {

	Warning(filename)
	Warning(object_name)
	Warning(object_owner)
	// Checking if the file name exists

	return "Valid"
}

func Invalid_op() {
	Break("Invalid option or number of arguments given run encore -h for help")
}

func WriteToFile(data string, location string, append string) {
	// preping data
	var d = []byte(data)
	// checking if file exists
	if append == "write" {
		if Existence(location) == true {
			// default is to overwrite the file
			if DeleteFile(location) == true { // file was deleted

				file, err := os.OpenFile(location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0400)
				if err != nil {
					Handle_err(err, "break")
				}
				if _, err := file.Write(d); err != nil {
					file.Close() // ignore error; Write error takes precedence
					Handle_err(err, "break")
				}
				if err := file.Close(); err != nil {
					Handle_err(err, "break")
				}
			} else {
				Break("Im lazy the file could not be deleted")
			}

		} else {
			// Nothing needs to be deleted just write it
			file, err := os.OpenFile(location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0400)
			if err != nil {
				Handle_err(err, "break")
			}
			if _, err := file.Write(d); err != nil {
				file.Close() // ignore error; Write error takes precedence
				Handle_err(err, "break")
			}
			if err := file.Close(); err != nil {
				Break("File stream incorrectly terminated")
			}
		}

	} else if append == "append" {
		// Nothing needs to be deleted just write it
		file, err := os.OpenFile(location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0400)
		if err != nil {
			Handle_err(err, "break")
		}
		if _, err := file.Write(d); err != nil {
			file.Close() // ignore error; Write error takes precedence
			Handle_err(err, "break")
		}
		if err := file.Close(); err != nil {
			Break("File stream incorrectly terminated")
		}

	} else {
		Warning("Invalid option given")
	}

}

func DeleteFile(filename string) bool {
	if Existence(filename) == true {
		del := os.Remove(filename)
		if del != nil {
			Handle_err(del, "warn")
			return false
		}
		return true

	} else {
		var msg string = "File not found : "
		msg += string("'" + filename + "'")
		Warning(msg)
		return true
	}
}

func Existence(filename string) bool {
	_, error := os.Stat(filename)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		return false
	} else {
		return true
	}
}

// Thanks https://mrwaggel.be/post/generate-md5-hash-of-a-file-in-golang
func Hash_file_md5(filePath string) (string, error) {
	var returnMD5String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}
	defer file.Close()
	var hash = md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	var hashInBytes = hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

func Handle_err(msg error, action string) {

	var error_message string = msg.Error()
	if action == "break" {
		Break(error_message)
	} else if action == "warn" {
		Warning(error_message)
	} else {
		Warning(error_message)
	}

}
