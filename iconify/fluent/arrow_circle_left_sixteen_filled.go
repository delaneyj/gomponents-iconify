package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleLeftSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14ZM4.538 8.191c.024.058.06.113.106.16l.003.003l2.5 2.5a.5.5 0 0 0 .707-.708L6.207 8.5H11a.5.5 0 0 0 0-1H6.207l1.647-1.646a.5.5 0 1 0-.708-.708l-2.5 2.5l-.002.003a.498.498 0 0 0-.106.542Z"/>`),
		g.Group(children),
	)
}