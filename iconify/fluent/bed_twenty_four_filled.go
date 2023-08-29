package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.25 11a2.75 2.75 0 0 1 2.745 2.582l.005.168v6.5a.75.75 0 0 1-1.493.102l-.007-.102V18h-17v2.25a.75.75 0 0 1-.648.743L2.75 21a.75.75 0 0 1-.743-.648L2 20.25v-6.5a2.75 2.75 0 0 1 2.582-2.745L4.75 11h14.5ZM6.75 4h10.5a2.75 2.75 0 0 1 2.745 2.582L20 6.75V10h-3l-.007-.117a1 1 0 0 0-.876-.876L16 9h-2a1 1 0 0 0-.993.883L13 10h-2l-.007-.117a1 1 0 0 0-.876-.876L10 9H8a1 1 0 0 0-.993.883L7 10H4V6.75a2.75 2.75 0 0 1 2.582-2.745L6.75 4Z"/>`),
		g.Group(children),
	)
}