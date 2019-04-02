package content

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"gopkg.in/yaml.v2"
)

// ParseYAMLFrontMatter will convert an input byte array like ---<stuff>---\n<body> into v as YAML and <body> as return value
func ParseYAMLFrontMatter(b []byte, v interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(b)

	var insideFrontMatter bool
	var yamlStartIndex int
	var yamlEndIndex int

	for {
		line, err := buf.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if strings.TrimSpace(line) != "---" {
			continue
		}

		if !insideFrontMatter {
			insideFrontMatter = true
			yamlStartIndex = len(b) - buf.Len()
		} else {
			yamlEndIndex = len(b) - buf.Len()
			break
		}
	}

	// if we get to here and we're not inside front matter then the entire string is body
	if !insideFrontMatter {
		return b, nil
	}

	if insideFrontMatter && yamlEndIndex == 0 {
		return nil, fmt.Errorf("Unexplained front matter parser error; insideFrontMatter: %v, yamlStartIndex: %v, yamlEndIndex: %v", insideFrontMatter, yamlStartIndex, yamlEndIndex)
	}

	err := yaml.Unmarshal(b[yamlStartIndex:yamlEndIndex], v)

	if err != nil {
		return nil, err
	}

	return bytes.TrimSpace(b[yamlEndIndex:]), nil
}
