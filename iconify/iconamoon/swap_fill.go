package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.707 20.707a1 1 0 0 1-1.414 0l-4-4A1 1 0 0 1 13 15h3V4a1 1 0 1 1 2 0v11h3a1 1 0 0 1 .707 1.707l-4 4Zm-10-17.414a1 1 0 0 0-1.414 0l-4 4A1 1 0 0 0 3 9h3v11a1 1 0 1 0 2 0V9h3a1 1 0 0 0 .707-1.707l-4-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}