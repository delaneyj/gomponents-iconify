package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M13 26h2c2.76 0 5-2.24 5-5s-2.24-5-5-5h-4c-2.76 0-5 2.24-5 5v9a6 6 0 0 0 6 6h24a6 6 0 0 0 6-6v-4a4 4 0 0 0-4-4H24m10 22H14a2 2 0 0 1-2-2v-6h24v6c0 1.1-.9 2-2 2Z"/><path d="M6 20c-1 0-2-2-2-4v-4c0-4.42 3.58-8 8-8h24c4.42 0 8 3.58 8 8v4c0 3.31-2.69 6-6 6H19.9"/></g>`),
		g.Group(children),
	)
}