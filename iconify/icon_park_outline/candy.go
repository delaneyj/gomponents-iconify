package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Candy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><circle cx="24" cy="24" r="10" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" d="M24 28a4 4 0 0 1-4-4"/><path stroke-linejoin="round" d="m16.688 16.813l-12.78-1.846L14.842 4.033l1.845 12.78Zm14.625 14.5l12.779 1.845l-10.934 10.934l-1.846-12.78Z"/></g>`),
		g.Group(children),
	)
}