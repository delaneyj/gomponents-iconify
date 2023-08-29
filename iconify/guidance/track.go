package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Track(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15 21.5c-1 0-1.75-1.5-1.75-1.5c-.75-1.5-.75-2.5-.75-4L8 13.5v-.25L8.462 12m5.038-2V9A2.5 2.5 0 0 0 11 6.5h-1L6.909 8.266a.811.811 0 0 0 .075 1.447l5.766 2.537M5.5 20c0-1 1.5-1.75 1.5-1.75c1.327-.664 2.263-.74 3.5-.749M0 23.5h23.5v-.25s-5-9.75-5-13.25A3.5 3.5 0 0 1 22 6.5m-10.148-2s-1.6-1-1.6-2.25c0-.966.784-1.75 1.75-1.75c.967 0 1.746.784 1.746 1.75c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}