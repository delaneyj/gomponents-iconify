package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Import(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M5 9.982v10h14v-10h-4v-2h6v14H3v-14h6v2H5Z"/><path d="M13 2h-2v12.053l-2.535-2.536l-1.415 1.415l4.95 4.95l4.95-4.95l-1.414-1.415L13 14.053V2Z"/></g>`),
		g.Group(children),
	)
}