package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleChartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 16a3 3 0 1 1 0 6a3 3 0 0 1 0-6ZM6 12a4 4 0 1 1 0 8a4 4 0 0 1 0-8Zm8.5-10a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11Z"/>`),
		g.Group(children),
	)
}