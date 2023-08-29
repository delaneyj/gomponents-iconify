package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneXMobiledata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17V9H4V7h4v10H6Zm4.35 0l3.15-5.3L10.65 7H13l1.65 2.75L16.35 7h2.3l-2.8 4.7L19 17h-2.35l-2-3.35l-2 3.35h-2.3Z"/>`),
		g.Group(children),
	)
}