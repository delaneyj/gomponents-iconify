package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatCursorTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 1a6 6 0 0 0-5.27 8.872l-.71 2.49a.5.5 0 0 0 .638.612l2.338-.779C5.88 12.707 6.906 13 8 13c.34 0 .675-.028 1-.083v-4.17c0-1.49 1.743-2.298 2.88-1.335l1.801 1.524A6 6 0 0 0 8 1Zm2 7.746v9.003c0 .74.96 1.033 1.373.418l1.978-2.946a.5.5 0 0 1 .415-.221h3.487c.698 0 1.018-.871.484-1.322l-6.502-5.504A.75.75 0 0 0 10 8.746Z"/>`),
		g.Group(children),
	)
}