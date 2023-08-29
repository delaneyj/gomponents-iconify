package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShelterEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4 5v3h6v2H2V5.67L1 6V4l9-3v2L4 5z" fill="currentColor"/>`),
		g.Group(children),
	)
}