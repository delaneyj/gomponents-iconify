package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Question(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-dasharray="32" stroke-dashoffset="32" stroke-linecap="round" stroke-width="2" d="M7 8C7 5.23858 9.23857 3 12 3C14.7614 3 17 5.23858 17 8C17 9.6356 16.2147 11.0878 15.0005 12C14.1647 12.6279 12 14 12 17"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.5s" values="32;0"/></path><circle cx="12" cy="21" r="1" fill="currentColor" fill-opacity="0"><animate fill="freeze" attributeName="fill-opacity" begin="0.5s" dur="0.2s" values="0;1"/></circle>`),
		g.Group(children),
	)
}