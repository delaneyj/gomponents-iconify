package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LampOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.56 4l2.67 8H6.78l2.66-8h5.12M16 2H8L4 14h16L16 2m-5 13h2v5h5v2H6v-2h5v-5Z"/>`),
		g.Group(children),
	)
}