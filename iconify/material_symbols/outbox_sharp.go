package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutboxSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 14V9.85l-1.6 1.6L8 10l4-4l4 4l-1.4 1.45l-1.6-1.6V14h-2Zm-8 7V3h18v18H3Zm2-2h14v-3h-3q-.75.95-1.788 1.475T12 18q-1.175 0-2.212-.525T8 16H5v3Zm7-3q.95 0 1.725-.55T14.8 14H19V5H5v9h4.2q.3.9 1.075 1.45T12 16Zm-7 3v-3h3q.75.95 1.788 1.475T12 18q1.175 0 2.212-.525T16 16h3v3H5Z"/>`),
		g.Group(children),
	)
}