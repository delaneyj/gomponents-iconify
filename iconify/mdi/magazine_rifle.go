package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagazineRifle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 1v2h1v10l-3 6l9 4l4-10V3h1V1m-9 2h6v2h-4.12v8.45L9.6 18.14L8 17.5l2-4Z"/>`),
		g.Group(children),
	)
}