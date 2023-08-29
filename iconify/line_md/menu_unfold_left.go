package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuUnfoldLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="10" stroke-dashoffset="10" d="M21 9L18 12L21 15"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="10;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M19 5H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="16;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M14 12H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="16" stroke-dashoffset="16" d="M19 19H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="16;0"/></path></g>`),
		g.Group(children),
	)
}