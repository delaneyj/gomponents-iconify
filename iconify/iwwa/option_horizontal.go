package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OptionHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<g fill="currentColor"><path d="M31.013 23.112a3.113 3.113 0 1 1-.002-6.226a3.113 3.113 0 0 1 .002 6.226z"/><circle cx="20.329" cy="20" r="3.112"/><circle cx="9.643" cy="20" r="3.112"/></g>`),
		g.Group(children),
	)
}