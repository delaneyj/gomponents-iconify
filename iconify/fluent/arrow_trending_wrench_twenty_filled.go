package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTrendingWrenchTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.75 3a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-.342.63a1.383 1.383 0 0 0-1.15-.9L16 7.979V5.561l-4.22 4.22a.75.75 0 0 1-1.06 0L9 8.06l-4.72 4.72a.75.75 0 0 1-1.06-1.061l5.25-5.25a.75.75 0 0 1 1.06 0l1.72 1.72l3.69-3.69h-2.69a.75.75 0 0 1 0-1.5h4.5Zm-.648 6.712c.26-.26.155-.696-.21-.739a3.5 3.5 0 0 0-3.704 4.652L9.475 16.34a1.5 1.5 0 0 0 2.121 2.121l2.713-2.713a3.5 3.5 0 0 0 4.653-3.704c-.043-.365-.479-.47-.739-.21l-.97.97a1.5 1.5 0 1 1-2.121-2.121l.97-.97Z"/>`),
		g.Group(children),
	)
}