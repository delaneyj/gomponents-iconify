package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRenminbi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 9v8a2 2 0 1 0 4 0m0-8H5m14-4H5m4 4v4c0 2.5-.667 4-2 6"/>`),
		g.Group(children),
	)
}