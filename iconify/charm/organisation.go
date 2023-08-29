package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Organisation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 1.75h3.5v3.5h-3.5zm4 9h3.5v3.5h-3.5zm-8 0h3.5v3.5h-3.5zm5.75-5v2m-3.75 2.5v-2h7.5v2"/>`),
		g.Group(children),
	)
}