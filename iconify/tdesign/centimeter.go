package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Centimeter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 4h10a2 2 0 0 1 2 2v14h-2V6h-3v14h-2V6h-3v14h-2V4ZM1 6a2 2 0 0 1 2-2h7v2H3v12h7v2H3a2 2 0 0 1-2-2V6Z"/>`),
		g.Group(children),
	)
}