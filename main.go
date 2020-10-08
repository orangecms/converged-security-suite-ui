package main

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/9elements/converged-security-suite/pkg/hwapi"
)

func main() {
	var label string
	tpmTss, err := hwapi.NewTPM()
	if err != nil {
		label = fmt.Sprintf("Couldn't set up tpm connection: %v\n", err)
	}
	defer tpmTss.Close()

	switch tpmTss.Version {
	case hwapi.TPMVersion12:
		label = "TPM 1.2 found, not supported yet"
		break
	case hwapi.TPMVersion20:
		label = "TPM 2.0 found, lucky you ;)"
		break
	}

	a := app.New()
	win := a.NewWindow("Converged Security Suite")
	win.SetContent(widget.NewVBox(
		widget.NewLabel(label),
		widget.NewButton("Ok", func() {
			a.Quit()
		}),
	))
	win.ShowAndRun()
}
