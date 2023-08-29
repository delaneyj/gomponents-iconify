package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowURightDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m216 176l-48 48l-48-48Z" opacity=".2"/><path d="M223.39 172.94A8 8 0 0 0 216 168h-40V88a64 64 0 0 0-128 0v88a8 8 0 0 0 16 0V88a48 48 0 0 1 96 0v80h-40a8 8 0 0 0-5.66 13.66l48 48a8 8 0 0 0 11.32 0l48-48a8 8 0 0 0 1.73-8.72ZM168 212.69L139.31 184h57.38Z"/></g>`),
		g.Group(children),
	)
}