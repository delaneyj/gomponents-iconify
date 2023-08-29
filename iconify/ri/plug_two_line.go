package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 18v2h6v2h-6a2 2 0 0 1-2-2v-2H8a4 4 0 0 1-4-4V7a1 1 0 0 1 1-1h2V2h2v4h6V2h2v4h2a1 1 0 0 1 1 1v7a4 4 0 0 1-4 4h-3Zm-5-2h8a2 2 0 0 0 2-2v-3H6v3a2 2 0 0 0 2 2Zm10-8H6v1h12V8Zm-6 6.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2ZM11 2h2v3h-2V2Z"/>`),
		g.Group(children),
	)
}