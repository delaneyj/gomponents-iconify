package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="M33 5H5v38h28V5Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M33 21h10v22H33"/><path stroke-linecap="round" d="M13 21h12m-6-6v12"/></g>`),
		g.Group(children),
	)
}