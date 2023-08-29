package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsExchangeV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.984 15a1 1 0 0 0 1.848.53l2.688-2.687a1 1 0 0 0-1.415-1.414l-1.12 1.12V5a1 1 0 0 0-2 0v10Zm-1.968-6a1 1 0 0 0-1.848-.53l-2.687 2.687a1 1 0 1 0 1.414 1.414l1.121-1.12V19a1 1 0 1 0 2 0V9Z"/>`),
		g.Group(children),
	)
}