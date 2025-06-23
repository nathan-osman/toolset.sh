package server

import (
	"mime"
	"strings"
)

func parseAcceptHeader(v string) (outputFormat, error) {
	for _, m := range strings.Split(v, ",") {
		mediaType, _, err := mime.ParseMediaType(m)
		if err != nil {
			return 0, err
		}
		switch {
		case mediaType == "application/json":
			return outputJson, nil
		case mediaType == "text/plain":
			return outputText, nil
		}
	}
	return outputHtml, nil
}
