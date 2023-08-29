package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4.854 6.146L4.5 5.793l-.707.707l.353.354l.708-.708ZM7.5 9.5l-.354.354l.354.353l.354-.353L7.5 9.5Zm3.354-2.646l.353-.354l-.707-.707l-.354.353l.708.708Zm-6.708 0l3 3l.708-.708l-3-3l-.708.708Zm3.708 3l3-3l-.708-.708l-3 3l.708.708ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1Zm-1 0A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0Z"/>`),
		g.Group(children),
	)
}