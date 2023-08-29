package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ladder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 42V10c0-6 6-6 6-6m-6 10h18m-18 8h18m-18 8h18m-18 8h18m6 5V11c0-6 6-6 6-6"/>`),
		g.Group(children),
	)
}