package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M22 22H2"/><path d="M21 22v-7.5a1.5 1.5 0 0 0-1.5-1.5h-3a1.5 1.5 0 0 0-1.5 1.5V22" opacity=".5"/><path d="M15 22V5c0-1.414 0-2.121-.44-2.56C14.122 2 13.415 2 12 2c-1.414 0-2.121 0-2.56.44C9 2.878 9 3.585 9 5v17"/><path d="M9 22V9.5A1.5 1.5 0 0 0 7.5 8h-3A1.5 1.5 0 0 0 3 9.5V22" opacity=".5"/></g>`),
		g.Group(children),
	)
}