package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adjustment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 8a3 3 0 0 1 3-3h32a3 3 0 0 1 3 3v32a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V8Zm31 4L12 36m0-20h8m7 17h8M16 12v8"/>`),
		g.Group(children),
	)
}