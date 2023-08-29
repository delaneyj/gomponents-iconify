package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kuula(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.388 0 0 5.388 0 12s5.388 12 12 12s12-5.388 12-12S18.612 0 12 0Zm0 2.547A9.433 9.433 0 0 1 21.453 12A9.433 9.433 0 0 1 12 21.453A9.433 9.433 0 0 1 2.547 12A9.433 9.433 0 0 1 12 2.547Zm-.606 5.366l4.372-.693l2.01 3.944l-3.13 3.13l-3.944-2.01z"/>`),
		g.Group(children),
	)
}