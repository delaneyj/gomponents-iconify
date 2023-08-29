package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyKztThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 96a4 4 0 0 1-4 4h-68v116a4 4 0 0 1-8 0V100H56a4 4 0 0 1 0-8h144a4 4 0 0 1 4 4ZM56 60h144a4 4 0 0 0 0-8H56a4 4 0 0 0 0 8Z"/>`),
		g.Group(children),
	)
}