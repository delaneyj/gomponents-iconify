package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronSmallLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path stroke="currentColor" stroke-dasharray="8" stroke-dashoffset="8" stroke-linecap="round" stroke-width="2" d="M9 12L14 7M9 12L14 17" fill="currentColor"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.3s" values="8;0"/></path>`),
		g.Group(children),
	)
}