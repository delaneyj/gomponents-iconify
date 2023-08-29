package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m8.146 9.146l-.353.354l.707.707l.354-.353l-.708-.708ZM10.5 7.5l.354.354l.353-.354l-.353-.354l-.354.354ZM8.854 5.146L8.5 4.793l-.707.707l.353.354l.708-.708Zm0 4.708l2-2l-.708-.708l-2 2l.708.708Zm2-2.708l-2-2l-.708.708l2 2l.708-.708ZM10.5 7H4v1h6.5V7Z"/>`),
		g.Group(children),
	)
}