package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m5.033 14.828l1.415 1.415L10.69 12L6.448 7.757L5.033 9.172L7.862 12l-2.829 2.828ZM15 14h-4v2h4v-2Z"/><path fill-rule="evenodd" d="M2 2a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H2Zm20 2H2v16h20V4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}