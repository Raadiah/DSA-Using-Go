package theme

import "github.com/fatih/color"

func WriteAlert(message string) {
	color.Yellow(message)
}

func WriteInstruction(message string) {
	color.Cyan(message)
}

func SuccessMessage(message string, formatter ...any) {
	if len(formatter) > 0 {
		color.Green(message, formatter...)
	} else {
		color.Green(message)
	}
}

func ErrorMessage(message string) {
	color.Red(message)
}
