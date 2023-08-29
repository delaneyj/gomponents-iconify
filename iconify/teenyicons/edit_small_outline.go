package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m4.5 8.5l-.354-.354L4 8.293V8.5h.5Zm4-4l.354-.354a.5.5 0 0 0-.708 0L8.5 4.5Zm2 2l.354.354a.5.5 0 0 0 0-.708L10.5 6.5Zm-4 4v.5h.207l.147-.146L6.5 10.5Zm-2 0H4a.5.5 0 0 0 .5.5v-.5Zm.354-1.646l4-4l-.708-.708l-4 4l.708.708Zm3.292-4l2 2l.708-.708l-2-2l-.708.708Zm2 1.292l-4 4l.708.708l4-4l-.708-.708ZM6.5 10h-2v1h2v-1Zm-1.5.5v-2H4v2h1Z"/>`),
		g.Group(children),
	)
}