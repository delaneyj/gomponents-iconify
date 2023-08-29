package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CPlusPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12.207 16.278a6 6 0 1 1 .071-8.485l1.414-1.414a8 8 0 1 0-.071 11.314l-1.414-1.415Z"/><path d="M15 9h-2v2h-2v2h2v2h2v-2h2v-2h-2V9Zm5 0h2v2h2v2h-2v2h-2v-2h-2v-2h2V9Z"/></g>`),
		g.Group(children),
	)
}