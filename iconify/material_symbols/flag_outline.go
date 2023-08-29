package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21V4h9l.4 2H20v10h-7l-.4-2H7v7H5Zm7.5-11Zm2.15 4H18V8h-5.25l-.4-2H7v6h7.25l.4 2Z"/>`),
		g.Group(children),
	)
}