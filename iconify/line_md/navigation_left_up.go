package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M21 15H12C9.23858 15 7 12.7614 7 10V4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M7 3L3 7M7 3L11 7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></path></g>`),
		g.Group(children),
	)
}