package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.143 2.153a1 1 0 0 1 1.715 0l3.999 6.665a1 1 0 0 1-.858 1.515H2.001a1 1 0 0 1-.858-1.515l4-6.665z"/>`),
		g.Group(children),
	)
}