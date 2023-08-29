package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m9.5 17l-5-5l5-5"/><path d="M4.5 12h10c1.667 0 5-1 5-5" opacity=".5"/></g>`),
		g.Group(children),
	)
}