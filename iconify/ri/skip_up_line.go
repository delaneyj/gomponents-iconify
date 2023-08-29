package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipUpLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 13.914l4.793 4.793l1.414-1.414L12 11.086l-6.207 6.207l1.414 1.414L12 13.914ZM6 7h12v2H6V7Z"/>`),
		g.Group(children),
	)
}