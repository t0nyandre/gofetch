package internal

import (
	"bufio"
	"bytes"
	"errors"
	"os"
	"strings"
)

func ReadOSRelease() (osrelease map[string]string, err error) {
	osrelease = make(map[string]string)

	lines, err := parseFile("/etc/os-release")
	if err != nil {
		return nil, err
	}

	for _, v := range lines {
		key, value, err := parseLine(v)
		if err != nil {
			return nil, err
		}
		osrelease[key] = value
	}
	return osrelease, nil
}

func parseFile(filename string) (lines []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseString(content string) (lines []string, err error) {
	in := bytes.NewBufferString(content)
	reader := bufio.NewReader(in)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func parseLine(line string) (key string, value string, err error) {
	err = nil

	if len(line) == 0 {
		err = errors.New("Skipping: Wont parse empty lines")
	}

	if line[0] == '#' {
		err = errors.New("Skipping: Comments")
	}

	splitString := strings.SplitN(line, "=", 2)
	if len(splitString) != 2 {
		err = errors.New("Can not extract key=value pair")
	}

	key = strings.Trim(splitString[0], " ")
	value = strings.Trim(splitString[1], " ")

	// Handle double quotes
	if strings.ContainsAny(value, `"`) {
		first := string(value[0:1])
		last := string(value[len(value)-1:])

		if first == last && strings.ContainsAny(first, `"'`) {
			value = strings.TrimPrefix(value, `'`)
			value = strings.TrimPrefix(value, `"`)
			value = strings.TrimSuffix(value, `'`)
			value = strings.TrimSuffix(value, `"`)
		}
	}

	value = strings.Replace(value, `\"`, `"`, -1)
	value = strings.Replace(value, `\$`, `$`, -1)
	value = strings.Replace(value, `\\`, `\`, -1)
	value = strings.Replace(value, "\\`", "`", -1)
	return
}
