package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.707 2.293a1 1 0 0 0-1.414 0l-2 2A1 1 0 0 0 10 6h1v5H6v-1a1 1 0 0 0-1.707-.707l-2 2a1 1 0 0 0 0 1.414l2 2A1 1 0 0 0 6 14v-1h5v5h-1a1 1 0 0 0-.707 1.707l2 2a1 1 0 0 0 1.414 0l2-2A1 1 0 0 0 14 18h-1v-5h5v1a1 1 0 0 0 1.707.707l2-2a1 1 0 0 0 0-1.414l-2-2A1 1 0 0 0 18 10v1h-5V6h1a1 1 0 0 0 .707-1.707l-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}