package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalTwelveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.5 1A1.5 1.5 0 0 1 10 2.5V4H2V2.5A1.5 1.5 0 0 1 3.5 1h5Zm2 4a.5.5 0 0 1 0 1h-9a.5.5 0 0 1 0-1h9ZM2 7v1.5A1.5 1.5 0 0 0 3.5 10h5A1.5 1.5 0 0 0 10 8.5V7H2Z"/>`),
		g.Group(children),
	)
}