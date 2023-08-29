package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArtTrackOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 19H1V5h14v14ZM3 17h10V7H3v10Zm14 2V5h2v14h-2Zm4 0V5h2v14h-2ZM4 15h8l-2.6-3.5L7.5 14l-1.4-1.85L4 15ZM3 7v10V7Z"/>`),
		g.Group(children),
	)
}