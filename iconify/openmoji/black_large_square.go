package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackLargeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M61 11.042H11v50h50v-50Z"/><path fill="#3F3F3F" d="M61 11.042H11v50h50v-50Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M61 11.042H11v50h50v-50Z"/>`),
		g.Group(children),
	)
}