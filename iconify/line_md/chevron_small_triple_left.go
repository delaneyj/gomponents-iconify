package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronSmallTripleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 12L19 7M14 12L19 17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="8;0"/></path><path d="M9 12L14 7M9 12L14 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="8;0"/></path><path d="M4 12L9 7M4 12L9 17"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.3s" values="8;0"/></path></g>`),
		g.Group(children),
	)
}