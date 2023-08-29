package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletBifold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3H7C4.79 3 3 4.79 3 7v10c0 2.21 1.79 4 4 4h12c1.11 0 2-.89 2-2V9a2 2 0 0 0-2-2V5a2 2 0 0 0-2-2m0 2v2H7c-.73 0-1.41.2-2 .54V7c0-1.1.9-2 2-2m8.5 10.5c-.83 0-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5s1.5.67 1.5 1.5s-.67 1.5-1.5 1.5Z"/>`),
		g.Group(children),
	)
}