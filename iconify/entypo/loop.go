package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Loop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20 7v7c0 1.103-.896 2-2 2H2c-1.104 0-2-.897-2-2V7a2 2 0 0 1 2-2h7V3l4 3.5L9 10V8H3v5h14V8h-3V5h4a2 2 0 0 1 2 2z"/>`),
		g.Group(children),
	)
}