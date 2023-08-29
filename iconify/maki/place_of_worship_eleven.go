package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaceOfWorshipEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.52 0L4 2v1h3V2L5.52 0zM4 4L2 5v5h7V5L7 4H4zm7 1.5V10h-1V5.5a.5.5 0 0 1 1 0zm-10 0V10H0V5.5a.5.5 0 0 1 1 0z" fill="currentColor"/>`),
		g.Group(children),
	)
}