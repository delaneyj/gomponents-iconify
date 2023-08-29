package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempColdFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 10.255V5a4 4 0 1 1 8 0v5.255a7 7 0 1 1-8 0ZM8 16a4 4 0 0 0 8 0H8Z"/>`),
		g.Group(children),
	)
}