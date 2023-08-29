package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m8.854 4.854l.353-.354l-.707-.707l-.354.353l.708.708ZM5.5 7.5l-.354-.354l-.353.354l.353.354L5.5 7.5Zm2.646 3.354l.354.353l.707-.707l-.353-.354l-.708.708Zm0-6.708l-3 3l.708.708l3-3l-.708-.708Zm-3 3.708l3 3l.708-.708l-3-3l-.708.708ZM7.5 14A6.5 6.5 0 0 1 1 7.5H0A7.5 7.5 0 0 0 7.5 15v-1ZM14 7.5A6.5 6.5 0 0 1 7.5 14v1A7.5 7.5 0 0 0 15 7.5h-1ZM7.5 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Zm0-1A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Z"/>`),
		g.Group(children),
	)
}