package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M12 12L3.5 20.5M12 12L20.5 3.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="14;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M21 3H13M3 21V13M21 3V11M3 21H11"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.2s" values="8;0"/></path></g>`),
		g.Group(children),
	)
}