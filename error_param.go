package main

type ErrorParam map[string]string

func NewErrorParam(err error) ErrorParam {
	return ErrorParam{"error": err.Error()}
}
