package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileArrowLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 13c.34 0 .67.04 1 .09V8l-6-6H6c-1.11 0-2 .89-2 2v16c0 1.11.89 2 2 2h7.81c-.51-.88-.81-1.9-.81-3c0-3.31 2.69-6 6-6m-6-9.5L18.5 9H13V3.5m7 16V18h-4v-2h4v-1.5l3 2.5l-3 2.5m-2 .5h4v2h-4v1.5L15 21l3-2.5V20Z"/>`),
		g.Group(children),
	)
}