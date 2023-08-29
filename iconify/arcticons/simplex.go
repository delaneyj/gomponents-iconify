package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Simplex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.75 42.5L42.5 14.75m-18.5 0l9.25-9.25M5.5 33.25l16-16m2.5 16l9.25 9.25M5.5 14.75l16 16M14.75 5.5L42.5 33.25"/>`),
		g.Group(children),
	)
}