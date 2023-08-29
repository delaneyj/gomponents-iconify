package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConciergeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 13h4V2H1v11Zm6 0h1.975L17 10V8h-4l-1.75.675l-.35-.925L13 7h9V5l-8-3l-7 1.975V13Zm3 9h13v-2H10v2Zm1-3h11q0-2.025-1.288-3.538T17.5 13.6v-1.625h-2V13.6q-1.95.35-3.225 1.863T11 19Z"/>`),
		g.Group(children),
	)
}