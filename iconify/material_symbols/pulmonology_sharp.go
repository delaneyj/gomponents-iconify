package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PulmonologySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 2h2v7.6l3 3l1.05-1.05L14.525 9V5H18.7l3.3 8.825V21h-9v-5l.4-3.2l-1.4-1.4l-1.4 1.4l.4 3.2v5H2v-7.175L5.3 5h4.2v4l-2.55 2.55L8 12.6l3-3V2Z"/>`),
		g.Group(children),
	)
}