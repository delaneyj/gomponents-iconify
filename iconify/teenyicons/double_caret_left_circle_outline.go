package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretLeftCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m6.146 9.854l.354.353l.707-.707l-.353-.354l-.708.708ZM4.5 7.5l-.354-.354l-.353.354l.353.354L4.5 7.5Zm2.354-1.646l.353-.354l-.707-.707l-.354.353l.708.708Zm2.292 4l.354.353l.707-.707l-.353-.354l-.708.708ZM7.5 7.5l-.354-.354l-.353.354l.353.354L7.5 7.5Zm2.354-1.646l.353-.354l-.707-.707l-.354.353l.708.708Zm-3 3.292l-2-2l-.708.708l2 2l.708-.708Zm-2-1.292l2-2l-.708-.708l-2 2l.708.708Zm5 1.292l-2-2l-.708.708l2 2l.708-.708Zm-2-1.292l2-2l-.708-.708l-2 2l.708.708ZM7.5 15A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1ZM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0ZM7.5 0A7.5 7.5 0 0 0 0 7.5h1A6.5 6.5 0 0 1 7.5 1V0Zm0 1A6.5 6.5 0 0 1 14 7.5h1A7.5 7.5 0 0 0 7.5 0v1Z"/>`),
		g.Group(children),
	)
}