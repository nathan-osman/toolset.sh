package server

import (
	"mime"
	"strings"
)

func parseAcceptHeader(v string) outputType {
	for _, m := range strings.Split(v, ",") {
		mediaType, _, err := mime.ParseMediaType(m)
		if err != nil {
			panic(err)
		}
		switch {
		case mediaType == "application/json":
			return outputJson
		case mediaType == "text/plain":
			return outputText
		}
	}
	return outputHtml
}
