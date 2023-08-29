package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterFlashLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.005 3.103l-4.95 4.95a7 7 0 1 0 9.9 0l-4.95-4.95Zm0-2.828l6.364 6.364a9 9 0 1 1-12.728 0L12.005.275Zm1 10.728h2.5l-4.5 6.5v-4.5h-2.5l4.5-6.5v4.5Z"/>`),
		g.Group(children),
	)
}