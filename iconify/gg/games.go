package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Games(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M15.47 11.293a1 1 0 1 0-1.415 1.414a1 1 0 0 0 1.415-1.414Zm.707-2.121a1 1 0 1 1 1.414 1.414a1 1 0 0 1-1.414-1.414Zm3.535 2.121a1 1 0 1 0-1.414 1.414a1 1 0 0 0 1.414-1.414Zm-3.535 2.121a1 1 0 1 1 1.414 1.415a1 1 0 0 1-1.414-1.415ZM6 13H4v-2h2V9h2v2h2v2H8v2H6v-2Z"/><path fill-rule="evenodd" d="M7 5a7 7 0 0 0 0 14h10a7 7 0 1 0 0-14H7Zm10 2H7a5 5 0 0 0 0 10h10a5 5 0 0 0 0-10Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}