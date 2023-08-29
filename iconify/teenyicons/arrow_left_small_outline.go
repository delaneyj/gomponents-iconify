package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m6.146 9.854l.354.353l.707-.707l-.353-.354l-.708.708ZM4.5 7.5l-.354-.354l-.353.354l.353.354L4.5 7.5Zm2.354-1.646l.353-.354l-.707-.707l-.354.353l.708.708Zm0 3.292l-2-2l-.708.708l2 2l.708-.708Zm-2-1.292l2-2l-.708-.708l-2 2l.708.708ZM4.5 8H11V7H4.5v1Z"/>`),
		g.Group(children),
	)
}