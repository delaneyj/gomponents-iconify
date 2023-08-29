package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeitOnline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.189 33.816L34.9 43.338L8.81 43.5L32.305 4.837M10.151 13.2l3.88-8.538l24.73-.162l-23.495 38.663"/>`),
		g.Group(children),
	)
}