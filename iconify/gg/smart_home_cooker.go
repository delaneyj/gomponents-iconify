package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartHomeCooker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15 16a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path d="M15 1H9v2h2v2H7a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4V9a4 4 0 0 0-4-4h-4V3h2V1Zm2 6H7a2 2 0 0 0-2 2h14a2 2 0 0 0-2-2Zm2 4H5v8a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-8Z"/></g>`),
		g.Group(children),
	)
}