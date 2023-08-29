package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-dasharray="20" stroke-dashoffset="20" stroke-linecap="round" stroke-width="2"><path d="M4 9H21"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="20;0"/></path><path d="M3 15H20"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="20;0"/></path><path d="M10 3L8 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.5s" dur="0.2s" values="20;0"/></path><path d="M16 3L14 21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.2s" values="20;0"/></path></g>`),
		g.Group(children),
	)
}