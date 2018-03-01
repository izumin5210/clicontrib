package clog

type logMode int

const (
	logModeNop logMode = iota
	logModeVerbose
	logModeDebug
)
