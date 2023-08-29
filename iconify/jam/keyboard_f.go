package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 0h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3zm0 6v2h2V6H3zm0-3v2h2V3H3zm0 6v2h2V9H3zm3 0v2h8V9H6zm0-3v2h2V6H6zm0-3v2h2V3H6zm3 3v2h2V6H9zm0-3v2h2V3H9zm6 6v2h2V9h-2zm-3-3v2h2V6h-2zm0-3v2h2V3h-2zm3 0v5h2V3h-2z"/>`),
		g.Group(children),
	)
}