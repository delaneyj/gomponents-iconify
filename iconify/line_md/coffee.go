package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="48" stroke-dashoffset="48" d="M17 4v9a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V4z"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.6s" values="48;0"/></path><path stroke-dasharray="14" stroke-dashoffset="14" d="M17 9H20C20.55 9 21 8.55 21 8V5C21 4.45 20.55 4 20 4H17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="1s" dur="0.2s" values="14;28"/></path><path stroke-dasharray="10" stroke-dashoffset="10" d="M11 20h8M11 20h-8"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="10;0"/></path></g>`),
		g.Group(children),
	)
}