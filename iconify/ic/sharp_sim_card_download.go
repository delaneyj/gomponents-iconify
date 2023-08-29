package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSimCardDownload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H10L4 8v14h16V2zm-8 15l-4-4h3V9.02L13 9v4h3l-4 4z"/>`),
		g.Group(children),
	)
}