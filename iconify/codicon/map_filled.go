package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 5.5V13l3.5-2.188v-7.5L2 5.5Zm7.5 7.188v-7.5l-3-1.875v7.5l3 1.874Zm1 0v-7.5L14 3v7.5l-3.5 2.188Z"/>`),
		g.Group(children),
	)
}