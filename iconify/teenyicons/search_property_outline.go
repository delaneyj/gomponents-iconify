package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchPropertyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M4.5 8.5H4a.5.5 0 0 0 .5.5v-.5Zm4 0V9a.5.5 0 0 0 .5-.5h-.5Zm0-2.5H9a.5.5 0 0 0-.146-.354L8.5 6Zm-2-2l.354-.354a.5.5 0 0 0-.708 0L6.5 4Zm-2 2l-.354-.354A.5.5 0 0 0 4 6h.5Zm10.354 8.146l-4-4l-.708.708l4 4l.708-.708ZM6.5 12A5.5 5.5 0 0 1 1 6.5H0A6.5 6.5 0 0 0 6.5 13v-1ZM12 6.5A5.5 5.5 0 0 1 6.5 12v1A6.5 6.5 0 0 0 13 6.5h-1ZM6.5 1A5.5 5.5 0 0 1 12 6.5h1A6.5 6.5 0 0 0 6.5 0v1Zm0-1A6.5 6.5 0 0 0 0 6.5h1A5.5 5.5 0 0 1 6.5 1V0Zm-2 9h4V8h-4v1ZM9 8.5V6H8v2.5h1Zm-.146-2.854l-2-2l-.708.708l2 2l.708-.708Zm-2.708-2l-2 2l.708.708l2-2l-.708-.708ZM4 6v2.5h1V6H4Z"/>`),
		g.Group(children),
	)
}