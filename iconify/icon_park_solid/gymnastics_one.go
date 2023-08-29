package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GymnasticsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGymnasticsOne0"><g fill="none" stroke="#fff" stroke-miterlimit="2" stroke-width="4"><path fill="#fff" d="M27 24a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m23 29l-2 7l-9-3l-4 11m13-8l1.49 6.48A2 2 0 0 0 24.43 44h10.58M7 23l16 6l-9-9l-1-9m-1-7h5c12 0 27 1.45 27 28"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGymnasticsOne0)"/>`),
		g.Group(children),
	)
}