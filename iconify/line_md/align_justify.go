package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width="2"><path d="M12 5H18M12 5H6"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.2s" values="8;0"/></path><path d="M12 10H18M12 10H6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.2s" dur="0.2s" values="8;0"/></path><path d="M12 15H18M12 15H6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.4s" dur="0.2s" values="8;0"/></path><path d="M12 20H18M12 20H6"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.2s" values="8;0"/></path></g>`),
		g.Group(children),
	)
}