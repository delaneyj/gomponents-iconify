package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m.5 10.5l-.354-.354l-.146.147v.207h.5Zm10-10l.354-.354a.5.5 0 0 0-.708 0L10.5.5Zm4 4l.354.354a.5.5 0 0 0 0-.708L14.5 4.5Zm-10 10v.5h.207l.147-.146L4.5 14.5Zm-4 0H0a.5.5 0 0 0 .5.5v-.5Zm.354-3.646l10-10l-.708-.708l-10 10l.708.708Zm9.292-10l4 4l.708-.708l-4-4l-.708.708Zm4 3.292l-10 10l.708.708l10-10l-.708-.708ZM4.5 14h-4v1h4v-1Zm-3.5.5v-4H0v4h1Z"/>`),
		g.Group(children),
	)
}