package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentDecreaseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m.5 7.5l-.354-.354a.5.5 0 0 0 0 .708L.5 7.5ZM3 4h12V3H3v1Zm4 4h8V7H7v1Zm-4 4h12v-1H3v1Zm-.146-2.854l-2-2l-.708.708l2 2l.708-.708Zm-2-1.292l2-2l-.708-.708l-2 2l.708.708Z"/>`),
		g.Group(children),
	)
}