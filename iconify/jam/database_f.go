package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 10a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3v-4a3 3 0 0 1 3-3zm5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm-3 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2zM5 6a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm3 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2z"/>`),
		g.Group(children),
	)
}