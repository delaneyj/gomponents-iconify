package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Beer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 6h-2V4a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v15c0 1.654 1.346 3 3 3h10c1.654 0 3-1.346 3-3v-1h2c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2zM8 17H6V7h2v10zm6 0h-2V7h2v10zm6-1h-2V8h2v8z"/>`),
		g.Group(children),
	)
}