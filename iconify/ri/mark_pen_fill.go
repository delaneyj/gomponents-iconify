package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkPenFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.95 2.39l5.657 5.657a1 1 0 0 1 0 1.414l-7.778 7.778l-2.122.707l-1.414 1.415a1 1 0 0 1-1.414 0l-4.243-4.243a1 1 0 0 1 0-1.414L6.05 12.29l.707-2.122l7.779-7.778a1 1 0 0 1 1.414 0Zm.707 3.536l-6.364 6.364l1.414 1.414l6.364-6.364l-1.414-1.414ZM4.283 16.886l2.828 2.828l-1.414 1.414l-4.243-1.414l2.829-2.828Z"/>`),
		g.Group(children),
	)
}