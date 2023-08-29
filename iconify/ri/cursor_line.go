package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.387 13.498l2.553 7.014l-4.698 1.71l-2.553-7.014l-3.899 2.445l1.619-16.02l11.537 11.232l-4.559.633Zm-.01 5.819l-2.715-7.46l2.96-.411l-5.64-5.49l-.791 7.83l2.531-1.587l2.715 7.46l.94-.342Z"/>`),
		g.Group(children),
	)
}