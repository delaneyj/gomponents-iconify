package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenbaseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 2.76l10 3.236l-.9 9.468l-9.1 6.86l-9.1-6.864L2.01 6L12 2.76Zm0 .825l-8.81 2.85L12 20.793l8.806-14.358L12 3.585Z"/>`),
		g.Group(children),
	)
}