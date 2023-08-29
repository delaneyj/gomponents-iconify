package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoEntry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32c-.001 16.568 13.432 30 30 30s30.001-13.432 30-30c.001-16.568-13.432-30-30-30zM9 38V26h46v12H9z"/>`),
		g.Group(children),
	)
}