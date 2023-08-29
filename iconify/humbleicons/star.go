package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m12 2l2.939 5.955l6.572.955l-4.756 4.635l1.123 6.545L12 17l-5.878 3.09l1.123-6.545L2.489 8.91l6.572-.955L12 2Z"/>`),
		g.Group(children),
	)
}