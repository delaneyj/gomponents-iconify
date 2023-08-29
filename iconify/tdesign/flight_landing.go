package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightLanding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.302 1.203l3.716.995l2.566 8.97l4.958 1.329a2.5 2.5 0 1 1-1.294 4.83L1.553 12.584l.46-7.124l3.173.85l.642 2.243l3.335.894l.139-8.245Zm1.956 2.594l-.138 8.245l-6.904-1.85l-.412-1.44l-.15 2.325l16.112 4.318a.5.5 0 1 0 .259-.966l-6.053-1.622l-2.566-8.97l-.148-.04ZM2 19h20v2H2v-2Z"/>`),
		g.Group(children),
	)
}