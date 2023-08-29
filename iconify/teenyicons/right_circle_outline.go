package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m6.146 10.146l-.353.354l.707.707l.354-.353l-.708-.708ZM9.5 7.5l.354.354l.353-.354l-.353-.354L9.5 7.5ZM6.854 4.146L6.5 3.793l-.707.707l.353.354l.708-.708Zm0 6.708l3-3l-.708-.708l-3 3l.708.708Zm3-3.708l-3-3l-.708.708l3 3l.708-.708ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1Zm0 1A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1Z"/>`),
		g.Group(children),
	)
}