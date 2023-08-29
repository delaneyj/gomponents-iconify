package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StayHydrated(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37 30.5c0-12.011-13-26-13-26s-13 13.989-13 26a13 13 0 0 0 26 0Zm-13 .565V20.916m6.43 10.149H24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.518 30.378A9.482 9.482 0 0 0 24 39.86"/>`),
		g.Group(children),
	)
}