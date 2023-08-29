package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PointerHand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 16V9.5c0-.828-.641-1.5-1.48-1.5C18 8 17 8.3 17 9.5M8 14V5.52M20 16c0 4-3.134 6-7 6s-5.196-1-8.196-6l-1.571-2.605c-.536-.868-.107-1.994.881-2.314a1.657 1.657 0 0 1 1.818.552L8 14.033"/><path d="M14 11V7.5A1.5 1.5 0 0 1 15.5 6v0A1.5 1.5 0 0 1 17 7.5V11m-6 0V6.5A1.5 1.5 0 0 1 12.5 5v0A1.5 1.5 0 0 1 14 6.5V11m-6 0V2.5A1.5 1.5 0 0 1 9.5 1v0A1.5 1.5 0 0 1 11 2.5V11"/></g>`),
		g.Group(children),
	)
}