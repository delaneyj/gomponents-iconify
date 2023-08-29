package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.005 7V5h-16v14h16v-2h-8a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h8Zm-17-4h18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1Zm10 6v6h7V9h-7Zm2 2h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}