package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowPreviousSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.296 11.736a.75.75 0 1 1-1.092 1.028l-4-4.25a.748.748 0 0 1 0-1.027l4-4.25a.75.75 0 1 1 1.092 1.028L8.78 8l3.516 3.736ZM4.75 3a.75.75 0 0 0-.75.75v8.5a.75.75 0 0 0 1.5 0v-8.5A.75.75 0 0 0 4.75 3Z"/>`),
		g.Group(children),
	)
}