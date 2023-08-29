package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m10.146 10.146l-.353.354l.707.707l.354-.353l-.708-.708ZM13.5 7.5l.354.354l.353-.354l-.353-.354l-.354.354Zm-2.646-3.354l-.354-.353l-.707.707l.353.354l.708-.708Zm-6.708 6.708l.354.353l.707-.707l-.353-.354l-.708.708ZM1.5 7.5l-.354-.354l-.353.354l.353.354L1.5 7.5Zm3.354-2.646l.353-.354l-.707-.707l-.354.353l.708.708Zm6 6l3-3l-.708-.708l-3 3l.708.708Zm3-3.708l-3-3l-.708.708l3 3l.708-.708Zm-9 3l-3-3l-.708.708l3 3l.708-.708Zm-3-2.292l3-3l-.708-.708l-3 3l.708.708Zm6.153-6.436l-2 12l.986.164l2-12l-.986-.164Z"/>`),
		g.Group(children),
	)
}