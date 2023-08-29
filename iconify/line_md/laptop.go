package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-dasharray="50" stroke-dashoffset="50" d="M12 17H5V7H19V17Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="50;0"/></path><path stroke-dasharray="20" stroke-dashoffset="20" d="M3 19H21"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.7s" dur="0.3s" values="20;0"/></path></g>`),
		g.Group(children),
	)
}