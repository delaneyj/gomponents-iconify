package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Remove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-dasharray="22" stroke-dashoffset="22" stroke-linecap="round" stroke-width="2"><path d="M19 5L5 19"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.3s" dur="0.3s" values="22;0"/></path><path d="M5 5L19 19"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="22;0"/></path></g>`),
		g.Group(children),
	)
}