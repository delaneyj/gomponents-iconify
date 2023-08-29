package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewComfortable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V8Zm3-1h14a1 1 0 0 1 1 1v3H4V8a1 1 0 0 1 1-1Zm-1 6v3a1 1 0 0 0 1 1h3v-4H4Zm6 4h9a1 1 0 0 0 1-1v-3H10v4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}