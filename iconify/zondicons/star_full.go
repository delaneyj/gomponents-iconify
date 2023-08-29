package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10 15l-5.878 3.09l1.123-6.545L.489 6.91l6.572-.955L10 0l2.939 5.955l6.572.955l-4.756 4.635l1.123 6.545z"/>`),
		g.Group(children),
	)
}