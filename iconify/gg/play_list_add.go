package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayListAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 5h12v2H2V5Zm0 4h12v2H2V9Zm8 4H2v2h8v-2Z"/><path d="M16 9h2v4h4v2h-4v4h-2v-4h-4v-2h4V9Z"/></g>`),
		g.Group(children),
	)
}