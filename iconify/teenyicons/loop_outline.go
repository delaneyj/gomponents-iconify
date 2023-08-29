package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m14.5 3.5l.354.354a.5.5 0 0 0 0-.708L14.5 3.5Zm-14 8l-.354-.354a.5.5 0 0 0 0 .708L.5 11.5ZM11.146.854l3 3l.708-.708l-3-3l-.708.708Zm3 2.292l-3 3l.708.708l3-3l-.708-.708Zm-10.292 11l-3-3l-.708.708l3 3l.708-.708Zm-3-2.292l3-3l-.708-.708l-3 3l.708.708ZM.5 12h11v-1H.5v1ZM15 8.5V7h-1v1.5h1ZM11.5 12A3.5 3.5 0 0 0 15 8.5h-1a2.5 2.5 0 0 1-2.5 2.5v1Zm3-9h-11v1h11V3ZM0 6.5V8h1V6.5H0ZM3.5 3A3.5 3.5 0 0 0 0 6.5h1A2.5 2.5 0 0 1 3.5 4V3Z"/>`),
		g.Group(children),
	)
}