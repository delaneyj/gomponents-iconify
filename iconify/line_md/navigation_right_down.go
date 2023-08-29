package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NavigationRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="26" stroke-dashoffset="26" d="M3 9H12C14.76142 9 17 11.2386 17 14V20"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="26;0"/></path><g stroke-dasharray="8" stroke-dashoffset="8"><path d="M17 21L21 17"/><path d="M17 21L13 17"/><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></g></g>`),
		g.Group(children),
	)
}