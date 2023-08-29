package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudyFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20.985a8.5 8.5 0 1 1 7.715-12.982A6.5 6.5 0 0 1 17 20.981V21H9v-.015Z"/>`),
		g.Group(children),
	)
}