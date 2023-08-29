package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstanceLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 7.653v8.694l7.5 4.342l7.5-4.342V7.653L12 3.311L4.5 7.653ZM12 1l9.5 5.5v11L12 23l-9.5-5.5v-11L12 1ZM6.499 9.97L11 12.577v5.049h2v-5.049l4.501-2.605l-1.002-1.731L12 10.844L7.501 8.24L6.499 9.97Z"/>`),
		g.Group(children),
	)
}