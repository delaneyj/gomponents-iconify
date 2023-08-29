package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nbc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 10.359v27.282l15.103-4.384l6.333 8.769L43 37.64V10.36l-16.564 4.385l-6.333-8.77L5 10.36Z"/>`),
		g.Group(children),
	)
}