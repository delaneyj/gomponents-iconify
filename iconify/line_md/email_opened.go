package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmailOpened(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-dasharray="64" stroke-dashoffset="64" d="M3 8.06083C3 7.71247 3.1813 7.38921 3.47855 7.20755L12 2L20.5214 7.20755C20.8187 7.38921 21 7.71247 21 8.06083V18C21 18.5523 20.5523 19 20 19H4C3.44772 19 3 18.5523 3 18V8.06083Z"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.6s" values="64;0"/></path><path stroke-dasharray="24" stroke-dashoffset="24" d="M3 8.5L12 14L21 8.5"><animate fill="freeze" attributeName="stroke-dashoffset" begin="0.6s" dur="0.4s" values="24;0"/></path></g>`),
		g.Group(children),
	)
}