package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M29.617 0H7v50h11V34h11.617c12.464 0 17.294-9.871 17.294-18.02C46.911 7.873 42.081 0 29.617 0zm-3.013 24H18V9h8.604c5.113 0 9.668 1.128 9.668 7.487C36.271 22.885 31.717 24 26.604 24z"/>`),
		g.Group(children),
	)
}