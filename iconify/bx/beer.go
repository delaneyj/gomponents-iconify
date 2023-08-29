package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6h-2V4a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v15c0 1.654 1.346 3 3 3h10c1.654 0 3-1.346 3-3v-1h2c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2zm-4 13c0 .551-.448 1-1 1H5c-.552 0-1-.449-1-1V5h12v14zm4-3h-2V8h2v8z"/><path fill="currentColor" d="M6 7h2v10H6zm6 0h2v10h-2z"/>`),
		g.Group(children),
	)
}