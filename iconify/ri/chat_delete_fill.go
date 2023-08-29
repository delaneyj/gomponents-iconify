package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatDeleteFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.455 19L2 22.5V4a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H6.455Zm6.96-8l2.474-2.475l-1.414-1.414L12 9.586L9.525 7.11L8.111 8.525L10.586 11L8.11 13.475l1.414 1.414L12 12.414l2.475 2.475l1.414-1.414L13.414 11Z"/>`),
		g.Group(children),
	)
}