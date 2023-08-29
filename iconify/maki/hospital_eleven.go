package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M10 4H7V1a1.08 1.08 0 0 0-1-1H5a1.08 1.08 0 0 0-1 1v3H1a1.08 1.08 0 0 0-1 1v1a1.08 1.08 0 0 0 1 1h3v3a1.08 1.08 0 0 0 1 1h1a1.08 1.08 0 0 0 1-1V7h3a1.08 1.08 0 0 0 1-1V5a1 1 0 0 0-1-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}