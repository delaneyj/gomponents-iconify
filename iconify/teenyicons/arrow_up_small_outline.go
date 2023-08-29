package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m5.146 6.146l-.353.354l.707.707l.354-.353l-.708-.708ZM7.5 4.5l.354-.354l-.354-.353l-.354.353l.354.354Zm1.646 2.354l.354.353l.707-.707l-.353-.354l-.708.708Zm-3.292 0l2-2l-.708-.708l-2 2l.708.708Zm1.292-2l2 2l.708-.708l-2-2l-.708.708ZM7 4.5V11h1V4.5H7Z"/>`),
		g.Group(children),
	)
}