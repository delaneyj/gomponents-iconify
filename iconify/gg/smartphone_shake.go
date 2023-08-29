package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneShake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13 14h-2v2h2v-2Z"/><path fill-rule="evenodd" d="M8 7a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V7Zm2 0h4v10h-4V7Z" clip-rule="evenodd"/><path d="M18 9h2v6h-2V9ZM0 14h2v-4H0v4Zm6 1H4V9h2v6Zm18-5h-2v4h2v-4Z"/></g>`),
		g.Group(children),
	)
}