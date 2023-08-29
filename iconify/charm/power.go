package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Power(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 1.75v6.5m4.25-5s2 1.298 2 4.75a6.25 6.25 0 1 1-12.5 0c0-3.452 2-4.75 2-4.75"/>`),
		g.Group(children),
	)
}