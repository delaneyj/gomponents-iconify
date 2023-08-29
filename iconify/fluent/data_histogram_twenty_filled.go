package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataHistogramTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 4.5A1.5 1.5 0 0 0 10.5 3h-1A1.5 1.5 0 0 0 8 4.5V17h4V4.5ZM13 6v11h3.5a.5.5 0 0 0 .5-.5V8a2 2 0 0 0-2-2h-2ZM5 9h2v8H3.5a.5.5 0 0 1-.5-.5V11a2 2 0 0 1 2-2Z"/>`),
		g.Group(children),
	)
}