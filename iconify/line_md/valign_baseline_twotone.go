package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ValignBaselineTwotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path fill="none" stroke-dasharray="20" stroke-dashoffset="20" d="M2.5 18.5H21.5M2.5 11.5H21.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path fill="currentColor" fill-opacity="0" stroke-dasharray="50" stroke-dashoffset="50" stroke-width="2" d="M12.75 6H18V18H6V6H12.75Z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.4s" values="50;0"/><animate fill="freeze" attributeName="fill-opacity" begin="0.7s" dur="0.15s" values="0;0.3"/></path></g>`),
		g.Group(children),
	)
}