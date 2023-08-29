package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditFlipH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-opacity=".5" d="M18 7a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-3v2h3a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3h-3v2h3Z"/><path d="M13 3h-2v18h2V3ZM5 8a1 1 0 0 1 1-1h3V5H6a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h3v-2H6a1 1 0 0 1-1-1V8Z"/></g>`),
		g.Group(children),
	)
}