package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotelService(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.5 10A1.5 1.5 0 0 1 7 8.5a5.5 5.5 0 0 1 11 0V19a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2v-2c0-1.38.71-2.444 1.88-3.175l4.424-2.765C14.359 10.4 15 9.744 15 8.5a2.5 2.5 0 1 0-5 0A1.5 1.5 0 0 1 8.5 10z"/>`),
		g.Group(children),
	)
}