package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func M(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M22.7 33.4c13.5-4.1 28.1 1.1 35.9 12.9l165.4 248L389.4 46.2c7.8-11.7 22.4-17 35.9-12.9S448 49.9 448 64v384c0 17.7-14.3 32-32 32s-32-14.3-32-32V169.7L250.6 369.8c-5.9 8.9-15.9 14.2-26.6 14.2s-20.7-5.3-26.6-14.2L64 169.7V448c0 17.7-14.3 32-32 32S0 465.7 0 448V64c0-14.1 9.2-26.5 22.7-30.6z"/>`),
		g.Group(children),
	)
}