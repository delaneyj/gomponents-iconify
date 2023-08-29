package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5.854 8.146L5.5 7.793l-.707.707l.353.354l.708-.708ZM7.5 10.5l-.354.354l.354.353l.354-.353L7.5 10.5Zm2.354-1.646l.353-.354l-.707-.707l-.354.353l.708.708ZM.5 7.5H0h.5Zm7-7V0v.5Zm0 14V14v.5Zm7-7H14h.5ZM5.146 8.854l2 2l.708-.708l-2-2l-.708.708Zm2.708 2l2-2l-.708-.708l-2 2l.708.708ZM8 10.5V4H7v6.5h1Zm-7-3A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1Zm1 0A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5h1Z"/>`),
		g.Group(children),
	)
}