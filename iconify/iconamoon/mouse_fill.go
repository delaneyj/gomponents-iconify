package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 9a7 7 0 0 1 14 0v6a7 7 0 1 1-14 0V9Zm8-2a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0V7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}