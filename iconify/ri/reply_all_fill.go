package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAllFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 4.5V9c5.523 0 10 4.477 10 10c0 .273-.01.543-.032.81a9.002 9.002 0 0 0-7.655-4.805L16 15h-2v4.5L6 12l8-7.5Zm-6 0v2.737L2.92 12l5.079 4.761L8 19.5L0 12l8-7.5Z"/>`),
		g.Group(children),
	)
}