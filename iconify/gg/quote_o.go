package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M20 5H4v14h16V5ZM4 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H4Z" clip-rule="evenodd"/><path d="M9.067 9.196h3l-2.134 5.608h-3l2.134-5.608Zm5 0h3l-2.134 5.608h-3l2.134-5.608Z"/></g>`),
		g.Group(children),
	)
}