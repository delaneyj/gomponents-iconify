package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m1 7l6.5 7L14 7M1 1.5l6.5 7l6.5-7"/>`),
		g.Group(children),
	)
}