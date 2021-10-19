package utils

import (
	"log"
	"net/url"
)

func IsUrl(str string) bool {
    _, err := url.ParseRequestURI(str)
    if err != nil {
        log.Fatalf(err.Error())
        return false
    }
	return true
}