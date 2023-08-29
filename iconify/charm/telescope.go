package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telescope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.75 5.75l1 2.5m3.5-4.5l1.5 3.5m-9 0l1 2.5l11.5-3.5l-2-4.5zm6 3.95v3m-3-.5L7 11.2l1.75-.5l2.5 3"/>`),
		g.Group(children),
	)
}