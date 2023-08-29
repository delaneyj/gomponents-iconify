package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 11a4.48 4.48 0 0 0 2.607-.832l3.613 3.612a.75.75 0 1 0 1.06-1.06l-3.612-3.613A4.5 4.5 0 1 0 6.5 11Zm-2-5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}