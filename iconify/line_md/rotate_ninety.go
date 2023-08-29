package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateNinety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="14" stroke-dashoffset="14" d="M12 6C15.3137 6 18 8.68629 18 12V14.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="14;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M18 15L21 12M18 15L15 12"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="6;0"/></path></g>`),
		g.Group(children),
	)
}