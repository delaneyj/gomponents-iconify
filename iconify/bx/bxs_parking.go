package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsParking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M13.5 3H5v18h4v-5h4.5c3.584 0 6.5-2.916 6.5-6.5S17.084 3 13.5 3zm0 9H9V7h4.5C14.879 7 16 8.121 16 9.5S14.879 12 13.5 12z" fill="currentColor"/>`),
		g.Group(children),
	)
}