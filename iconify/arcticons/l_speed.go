package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LSpeed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3.79 12.72l21.24 29.06l17.76-12.99l-4.87-6.66l-11.1 8.12l-16.37-22.4Z"/>`),
		g.Group(children),
	)
}