package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vectorlogozone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.458 0l-5.311 2.024l1.989.534l-4.847 16.085l-4.867-16.25H1.48L8.974 24h4.645l7.043-20.226l1.858.499Z"/>`),
		g.Group(children),
	)
}