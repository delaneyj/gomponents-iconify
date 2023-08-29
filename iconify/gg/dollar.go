package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dollar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 19v3h2v-3h1a4 4 0 0 0 0-8h-1V7h2v2h2V5h-4V2h-2v3h-1a4 4 0 1 0 0 8h1v4H9v-2H7v4h4Zm2-2h1a2 2 0 1 0 0-4h-1v4Zm-2-6V7h-1a2 2 0 1 0 0 4h1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}