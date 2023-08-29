package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="16" stroke-dashoffset="16" d="M8 18H7C4.5 18 3 16 3 14C3 12 4.5 10 7 10C8 10 8.5 10.5 8.5 10.5M16 18H17C19.5 18 21 16 21 14C21 12 19.5 10 17 10C16 10 15.5 10.5 15.5 10.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="16;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M17 10C17 10 17 9.5 17 9C17 6.5 15 4 12 4M7 10V9C7 6.5 9 4 12 4"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.3s" values="9;0"/></path><path stroke-dasharray="8" stroke-dashoffset="8" d="M12 20V14"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.9s" dur="0.2s" values="8;0"/></path><path stroke-dasharray="4" stroke-dashoffset="4" d="M12 13L14 15M12 13L10 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1.1s" dur="0.2s" values="4;0"/></path></g>`),
		g.Group(children),
	)
}