package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m9.146 6.854l.354.353l.707-.707l-.353-.354l-.708.708ZM7.5 4.5l.354-.354l-.354-.353l-.354.353l.354.354ZM5.146 6.146l-.353.354l.707.707l.354-.353l-.708-.708ZM14.5 7.5H14h.5Zm-7 7V14v.5Zm0-14V0v.5Zm-7 7H0h.5Zm9.354-1.354l-2-2l-.708.708l2 2l.708-.708Zm-2.708-2l-2 2l.708.708l2-2l-.708-.708ZM7 4.5V11h1V4.5H7Zm7 3A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1Zm-1 0A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0Z"/>`),
		g.Group(children),
	)
}