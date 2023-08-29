package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mitsubishi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="m17 19l7 11l7-11l-7-11l-7 11Zm13.667 21L24 30h13.333L44 40H30.667Zm-13.334 0L24 30H10.667L4 40h13.333Z"/>`),
		g.Group(children),
	)
}