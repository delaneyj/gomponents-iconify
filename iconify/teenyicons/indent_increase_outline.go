package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentIncreaseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m2.5 7.5l.354.354a.5.5 0 0 0 0-.708L2.5 7.5ZM3 4h12V3H3v1Zm4 4h8V7H7v1Zm-4 4h12v-1H3v1ZM.854 9.854l2-2l-.708-.708l-2 2l.708.708Zm2-2.708l-2-2l-.708.708l2 2l.708-.708Z"/>`),
		g.Group(children),
	)
}