package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandThreejs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 22L3 3l19 5.5z"/><path d="m12.573 17.58l-6.152-1.576l8.796-9.466l1.914 6.64"/><path d="M12.573 17.58L11 11l6.13 2.179M9.527 4.893L11 11L4.69 9.436z"/></g>`),
		g.Group(children),
	)
}