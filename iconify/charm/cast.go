package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M1.75 5.25v-2.5h12.5v10.5h-4.5m-8-5a5 5 0 0 1 5 5m-5-2.5a2.5 2.5 0 0 1 2.5 2.5m-2.5 0v0"/>`),
		g.Group(children),
	)
}