package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleArrowHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M12 12H3.5M12 12H20.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M3 12L7 16M21 12L17 16M3 12L7 8M21 12L17 8"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="6;0"/></path></g>`),
		g.Group(children),
	)
}