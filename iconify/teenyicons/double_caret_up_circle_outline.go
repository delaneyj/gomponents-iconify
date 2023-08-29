package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretUpCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m5.146 6.146l-.353.354l.707.707l.354-.353l-.708-.708ZM7.5 4.5l.354-.354l-.354-.353l-.354.353l.354.354Zm1.646 2.354l.354.353l.707-.707l-.353-.354l-.708.708Zm-4 2.292l-.353.354l.707.707l.354-.353l-.708-.708ZM7.5 7.5l.354-.354l-.354-.353l-.354.353l.354.354Zm1.646 2.354l.354.353l.707-.707l-.353-.354l-.708.708ZM.5 7.5H0h.5Zm7 7v.5v-.5Zm0-14V1V.5Zm7 7H14h.5Zm-8.646-.646l2-2l-.708-.708l-2 2l.708.708Zm1.292-2l2 2l.708-.708l-2-2l-.708.708Zm-1.292 5l2-2l-.708-.708l-2 2l.708.708Zm1.292-2l2 2l.708-.708l-2-2l-.708.708ZM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0ZM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0ZM15 7.5A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5h1Zm-1 0A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1Z"/>`),
		g.Group(children),
	)
}