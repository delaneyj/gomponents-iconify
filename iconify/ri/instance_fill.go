package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstanceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l9.5 5.5v11L12 23l-9.5-5.5v-11L12 1ZM6.499 9.97L11 12.578v5.048h2v-5.048l4.501-2.606l-1.002-1.731L12 10.845L7.501 8.24L6.499 9.97Z"/>`),
		g.Group(children),
	)
}