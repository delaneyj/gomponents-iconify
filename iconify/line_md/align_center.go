package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="9" stroke-dashoffset="9" d="M12 5H19M12 5H5"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="9;0"/></path><path stroke-dasharray="7" stroke-dashoffset="7" d="M12 10H17M12 10H7"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="7;0"/></path><path stroke-dasharray="11" stroke-dashoffset="11" d="M12 15H21M12 15H3"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="11;0"/></path><path stroke-dasharray="9" stroke-dashoffset="9" d="M12 20H19M12 20H5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="9;0"/></path></g>`),
		g.Group(children),
	)
}