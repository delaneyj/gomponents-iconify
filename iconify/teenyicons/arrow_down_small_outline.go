package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m9.854 8.854l.353-.354l-.707-.707l-.354.353l.708.708ZM7.5 10.5l-.354.354l.354.353l.354-.353L7.5 10.5ZM5.854 8.146L5.5 7.793l-.707.707l.353.354l.708-.708Zm3.292 0l-2 2l.708.708l2-2l-.708-.708Zm-1.292 2l-2-2l-.708.708l2 2l.708-.708ZM8 10.5V4H7v6.5h1Z"/>`),
		g.Group(children),
	)
}