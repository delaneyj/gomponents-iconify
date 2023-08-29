package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fancy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M242.275 0L66.038 98.243L0 341.512L101.902 512l173.115-95.18l65.787-245.08L242.274 0zm-64.35 164.126l-63.759-50.996l91.62-51.074l-27.861 102.07zm63.153-92.99l49.476 86.234l-77.202 15.344l27.726-101.579zm-149.16 70.857l75.553 60.43l-27.173 99.55l-95.774 14.608l47.394-174.588zm10.296 299.506l-53.056-88.765l80.643-12.302l-27.587 101.067zm62.651-91.153l62.242 51.235l-89.687 49.31l27.445-100.545zm84.31 22.208l-73.792-60.741l27.246-99.818l94.696-18.821l-48.15 179.38z"/>`),
		g.Group(children),
	)
}