package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m10.146 8.854l.354.353l.707-.707l-.353-.354l-.708.708ZM7.5 5.5l.354-.354l-.354-.353l-.354.353l.354.354ZM4.146 8.146l-.353.354l.707.707l.354-.353l-.708-.708Zm6.708 0l-3-3l-.708.708l3 3l.708-.708Zm-3.708-3l-3 3l.708.708l3-3l-.708-.708ZM1 7.5A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1Zm1 0A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5h1Z"/>`),
		g.Group(children),
	)
}