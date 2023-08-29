package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretDownCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m9.854 8.854l.353-.354l-.707-.707l-.354.353l.708.708ZM7.5 10.5l-.354.354l.354.353l.354-.353L7.5 10.5ZM5.854 8.146L5.5 7.793l-.707.707l.353.354l.708-.708Zm4-2.292l.353-.354l-.707-.707l-.354.353l.708.708ZM7.5 7.5l-.354.354l.354.353l.354-.353L7.5 7.5ZM5.854 5.146L5.5 4.793l-.707.707l.353.354l.708-.708ZM14.5 7.5H14h.5Zm-7-7V1V.5Zm0 14v.5v-.5Zm-7-7H0h.5Zm8.646.646l-2 2l.708.708l2-2l-.708-.708Zm-1.292 2l-2-2l-.708.708l2 2l.708-.708Zm1.292-5l-2 2l.708.708l2-2l-.708-.708Zm-1.292 2l-2-2l-.708.708l2 2l.708-.708ZM15 7.5A7.5 7.5 0 0 0 7.5 0v1A6.5 6.5 0 0 1 14 7.5h1ZM7.5 15A7.5 7.5 0 0 0 15 7.5h-1A6.5 6.5 0 0 1 7.5 14v1ZM0 7.5A7.5 7.5 0 0 0 7.5 15v-1A6.5 6.5 0 0 1 1 7.5H0Zm1 0A6.5 6.5 0 0 1 7.5 1V0A7.5 7.5 0 0 0 0 7.5h1Z"/>`),
		g.Group(children),
	)
}