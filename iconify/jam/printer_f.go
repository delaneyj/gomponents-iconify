package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 4h1a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3h-1V9H4v7H3a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h1v2h12V4zM6 20v-9h8v9H6zM6 4V0h8v4H6z"/>`),
		g.Group(children),
	)
}