package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 15.9a5.002 5.002 0 0 0 4-4.9V4H7v7a5.002 5.002 0 0 0 4 4.9V18H9v2h6v-2h-2v-2.1ZM9 6h6v5a3 3 0 1 1-6 0V6Z" clip-rule="evenodd"/><path d="M18 6h2v5h-2V6ZM6 6H4v5h2V6Z"/></g>`),
		g.Group(children),
	)
}