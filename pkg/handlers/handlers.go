package handlers

import "regexp"

var (
	redirectURLre = regexp.MustCompile(`^\/([\w-]+)$`)
	slugRe        = regexp.MustCompile(`^[\w-]+$`)
)
