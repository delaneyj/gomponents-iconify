package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monzo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M5.14 13.31L5 32l9.21 9.18v-18.9l9.86 9.86l9.37-9.38v18.72L43 31.93V13.21l-6.68-6.69l-12.25 12.24L11.9 6.6Z"/>`),
		g.Group(children),
	)
}