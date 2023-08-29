package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="15" stroke-dashoffset="15" d="M4 5H17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="15;0"/></path><path stroke-dasharray="12" stroke-dashoffset="12" d="M4 10H14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="12;0"/></path><path stroke-dasharray="18" stroke-dashoffset="18" d="M4 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="18;0"/></path><path stroke-dasharray="15" stroke-dashoffset="15" d="M4 20H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="15;0"/></path></g>`),
		g.Group(children),
	)
}