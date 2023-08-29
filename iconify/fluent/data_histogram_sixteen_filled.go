package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataHistogramSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 4a2 2 0 1 0-4 0v10h4V4ZM5 7H4a2 2 0 0 0-2 2v4.5a.5.5 0 0 0 .5.5H5V7Zm6 7h2.5a.5.5 0 0 0 .5-.5V7a2 2 0 0 0-2-2h-1v9Z"/>`),
		g.Group(children),
	)
}