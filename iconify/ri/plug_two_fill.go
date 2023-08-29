package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 18v2h6v2h-6a2 2 0 0 1-2-2v-2H8a4 4 0 0 1-4-4v-4h16v4a4 4 0 0 1-4 4h-3Zm4-12h2a1 1 0 0 1 1 1v2H4V7a1 1 0 0 1 1-1h2V2h2v4h6V2h2v4Zm-5 8.5a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM11 2h2v3h-2V2Z"/>`),
		g.Group(children),
	)
}