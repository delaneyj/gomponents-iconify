package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenFullFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 4.5a1 1 0 0 1 1-1h2a3 3 0 0 1 3 3v2a1 1 0 1 1-2 0v-2a1 1 0 0 0-1-1h-2a1 1 0 0 1-1-1Zm-11 4a1 1 0 1 0 2 0v-2a1 1 0 0 1 1-1h2a1 1 0 0 0 0-2h-2a3 3 0 0 0-3 3v2Zm17 7a1 1 0 1 0-2 0v2a1 1 0 0 1-1 1h-2a1 1 0 1 0 0 2h2a3 3 0 0 0 3-3v-2Zm-12 5a1 1 0 1 0 0-2h-2a1 1 0 0 1-1-1v-2a1 1 0 1 0-2 0v2a3 3 0 0 0 3 3h2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}