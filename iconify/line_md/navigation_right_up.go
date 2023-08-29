package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationRightUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M3 15H12C14.76142 15 17 12.7614 17 10V4"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8"><path d="M17 3L21 7"/><path d="M17 3L13 7"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></g></g>`),
		g.Group(children),
	)
}