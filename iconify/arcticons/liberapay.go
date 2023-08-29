package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Liberapay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.033 4.5c-1.366 5.694-4.736 20.82-5.306 23.384c-.484 2.181-1.179 6.511 2.872 7.342c4.252.872 21.277 3.932 23.82-9.785c2.306-12.447-13.625-9.5-13.625-9.5S18.281 38.94 17.41 43.5"/>`),
		g.Group(children),
	)
}