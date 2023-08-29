package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrZeroGramm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.12 7l17 17l-17 17l-4.62-4.67L17.83 24L5.5 11.67ZM24 41h18.5v-6.21H24Z"/>`),
		g.Group(children),
	)
}