package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOneOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m.5 9.5l-.354-.354L0 9.293V9.5h.5Zm9-9l.354-.354a.5.5 0 0 0-.708 0L9.5.5Zm5 5l.354.354a.5.5 0 0 0 0-.708L14.5 5.5Zm-9 9v.5h.207l.147-.146L5.5 14.5Zm-5 0H0a.5.5 0 0 0 .5.5v-.5Zm.354-4.646l9-9l-.708-.708l-9 9l.708.708Zm8.292-9l5 5l.708-.708l-5-5l-.708.708Zm5 4.292l-9 9l.708.708l9-9l-.708-.708ZM5.5 14h-5v1h5v-1Zm-4.5.5v-5H0v5h1ZM6.146 3.854l5 5l.708-.708l-5-5l-.708.708ZM8 15h7v-1H8v1Z"/>`),
		g.Group(children),
	)
}