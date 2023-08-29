package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CopperIud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M23 6h-4.132l1-2H23a1 1 0 0 1 1 1a1 1 0 0 1 1-1h3.132l1 2H25v10.257l-2-1V6ZM10 5a1 1 0 0 0 1 1h1.632l1-2H11a1 1 0 0 0-1 1Zm4.868 1l1-2h1.764l-1 2h-1.764Z"/><path fill-rule="evenodd" d="M20 40a4.002 4.002 0 0 1 3-3.874V25.493l2 1v9.633A4.002 4.002 0 0 1 24 44a4 4 0 0 1-4-4Zm4-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4Z" clip-rule="evenodd"/><path d="m25 20.257l-2-1v-1.764l2 1v1.764Zm0 2.236l-2-1v1.764l2 1v-1.764ZM38 5a1 1 0 0 0-1-1h-2.632l1 2H37a1 1 0 0 0 1-1Zm-6.632 1l-1-2h1.764l1 2h-1.764Z"/></g>`),
		g.Group(children),
	)
}