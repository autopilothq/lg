package lg

import (
	"regexp"
	"strings"
)

type regexpFilter struct {
	pattern *regexp.Regexp
}

func Regexp(pattern string) *regexpFilter {
	return &regexpFilter{
		regexp.MustCompile(pattern),
	}
}

func (f *regexpFilter) check(e *Entry) bool {
	return f.pattern.MatchString(e.Message)
}

type containsFilter struct {
	text string
}

func Contains(text string) *containsFilter {
	return &containsFilter{text}
}

func (f *containsFilter) check(e *Entry) bool {
	return strings.Contains(e.Message, f.text)
}

type exactLevelFilter struct {
	level Level
}

func AtLevel(level Level) *exactLevelFilter {
	return &exactLevelFilter{level}
}

func (f *exactLevelFilter) check(e *Entry) bool {
	return e.Level == f.level
}

type minLevelFilter struct {
	level Level
}

func AtLeastLevel(level Level) *minLevelFilter {
	return &minLevelFilter{level}
}

func (f *minLevelFilter) check(e *Entry) bool {
	return e.Level >= f.level
}

type maxLevelFilter struct {
	level Level
}

func AtMostLevel(level Level) *maxLevelFilter {
	return &maxLevelFilter{level}
}

func (f *maxLevelFilter) check(e *Entry) bool {
	return e.Level <= f.level
}
