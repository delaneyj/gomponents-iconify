package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ValignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 15.5H21.5M2.5 8.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 3H18V15H6V3H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/></path></g>`),
		g.Group(children),
	)
}