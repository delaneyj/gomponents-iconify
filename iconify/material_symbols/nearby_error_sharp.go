package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearbyErrorSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18v-8h2v8h-2Zm1 4q-.425 0-.713-.288T20 21q0-.425.288-.713T21 20q.425 0 .713.288T22 21q0 .425-.288.713T21 22Zm-9 .8L1.2 12L12 1.2l6 6v3.6l-6-6L4.8 12l7.2 7.2l6-6v3.6l-6 6Zm0-6.4L7.6 12L12 7.6l4.4 4.4l-4.4 4.4Z"/>`),
		g.Group(children),
	)
}