package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Marker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-dasharray="44" stroke-dashoffset="44" stroke-width="2" d="M6 14L17 3L21 7L10 18L6 14Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="44;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M9 17L6.5 19.5H2.5L7 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="20;0"/></path></g>`),
		g.Group(children),
	)
}