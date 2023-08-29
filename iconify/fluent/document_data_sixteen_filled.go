package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDataSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 5h3.25L9 1.25V4.5a.5.5 0 0 0 .5.5Zm0 1A1.5 1.5 0 0 1 8 4.5V1H4.5A1.5 1.5 0 0 0 3 2.5v11A1.5 1.5 0 0 0 4.5 15h7a1.5 1.5 0 0 0 1.5-1.5V6H9.5ZM6 12.5a.5.5 0 0 1-1 0v-6a.5.5 0 0 1 1 0v6Zm2.5 0a.5.5 0 0 1-1 0v-2a.5.5 0 0 1 1 0v2Zm2.5 0a.5.5 0 0 1-1 0v-4a.5.5 0 0 1 1 0v4Z"/>`),
		g.Group(children),
	)
}