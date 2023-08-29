package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingWindTurbine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 11a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="M10 11V8.427c0-.18.013-.358.04-.536l.716-4.828C10.82 2.466 11.353 2 12 2s1.18.466 1.244 1.063l.716 4.828c.027.178.04.357.04.536V11"/><path d="m13.01 9.28l2.235 1.276c.156.09.305.19.446.3l3.836 2.911c.487.352.624 1.04.3 1.596c-.325.556-1 .782-1.548.541l-4.555-1.68a3.624 3.624 0 0 1-.486-.231l-2.235-1.277"/><path d="m13 12.716l-2.236 1.277a3.624 3.624 0 0 1-.485.23l-4.555 1.681c-.551.241-1.223.015-1.548-.54c-.324-.557-.187-1.245.3-1.597l3.836-2.91a3.41 3.41 0 0 1 .446-.3l2.235-1.277M7 21h10m-7 0l1-7m2 0l1 7"/></g>`),
		g.Group(children),
	)
}