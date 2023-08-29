package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BubbleChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 16a3 3 0 1 1 0 6a3 3 0 0 1 0-6ZM6 12a4 4 0 1 1 0 8a4 4 0 0 1 0-8Zm10 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM6 14a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm8.5-12a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11Zm0 2a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Z"/>`),
		g.Group(children),
	)
}