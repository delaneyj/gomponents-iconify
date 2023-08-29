package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateOneHundredEighty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="24" stroke-dashoffset="24" d="M12 6C15.3137 6 18 8.68629 18 12C18 15.3137 15.3137 18 12 18H9.5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.4s" values="24;0"/></path><path stroke-dasharray="6" stroke-dashoffset="6" d="M9 18L12 21M9 18L12 15"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="6;0"/></path></g>`),
		g.Group(children),
	)
}