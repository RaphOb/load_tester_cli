package main

import (
	"errors"
	"flag"
	"math"
)

type CmdFlags struct {
	FileDescriptor string
	RateLimit      int
	Freq           string
	c              int
	nbR            int
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.FileDescriptor, "file", "spec.api", "File name is by default spec.apo. You can overwrite it with the option -file=")
	flag.IntVar(&cf.RateLimit, "rate", 0, "Number of rate desired, need to be coupled with freq")
	flag.IntVar(&cf.c, "c", math.MaxInt32, "Max concurrency, if not specified, unlimited")

	flag.IntVar(&cf.nbR, "r", 1, "Number total of requets.")
	flag.Func("freq", "You can choose frequency. It will fire rate by freq request. Allowed freq are: ms, s, min", func(flagValue string) error {
		for _, allowedValue := range []string{"ms", "s", "min"} {
			if flagValue == allowedValue {
				cf.Freq = flagValue
				return nil
			}
		}
		return errors.New(("must be one of 'ms, 's' or 'min'"))
	})

	flag.Parse()
	if (cf.RateLimit == 0) != (cf.Freq == "") {
		panic("both -rate and -freq must be set together, or neither")
	}
	return &cf
}
