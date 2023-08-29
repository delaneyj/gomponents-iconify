package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.005 7h1v10h-1v3a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v3Zm-2 10h-6a5 5 0 0 1 0-10h6V5h-16v14h16v-2Zm1-2V9h-7a3 3 0 1 0 0 6h7Zm-7-4h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}