package logs

import (
	"flag"
	"github.com/rs/zerolog"
)

const (
	TraceLvl   = "trace"
	DebugLvl   = "debug"
	InfoLvl    = "info"
	WarningLvl = "warn"
	ErrorLvl   = "error"
	FatalLvl   = "fatal"
)

func LogLvl(lvl string) {
	switch lvl {
	case TraceLvl:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case DebugLvl:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case InfoLvl:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case WarningLvl:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case ErrorLvl:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case FatalLvl:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}
}

func ParseLogLvl() {
	lvl := flag.String("LogLvl", InfoLvl, "Logging Lvl")
	flag.Parse()

	LogLvl(*lvl)
}
