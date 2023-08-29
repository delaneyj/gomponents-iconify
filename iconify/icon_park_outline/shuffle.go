package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shuffle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M30 42h12V30m0-12V6H30m12 0L6 42m25.5-10.5L34 34l5 5l2.5 2.5l.5.5M24 24L6 6l18 18Z"/>`),
		g.Group(children),
	)
}