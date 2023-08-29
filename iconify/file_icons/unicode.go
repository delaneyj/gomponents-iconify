package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unicode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M456.708 133.78c-24.5-14.126-24.5-49.63 0-63.754S512 73.653 512 101.903s-30.792 46.001-55.292 31.876zm-10.204 221.85V176.779h58.026v265.057l-78.294.016l-144.344-270.048v154.594c-4.79 66.46-47.996 122.478-145.729 120.481C49.495 446.473 7.657 399.895 0 327.217V87.232h64.678v208.075c0 65.35 23.343 94.786 75.974 94.786c54.065 0 75.58-26.204 75.58-94.786V87.232h84.644L446.504 355.63z"/>`),
		g.Group(children),
	)
}