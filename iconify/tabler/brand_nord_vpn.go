package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandNordVpn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9.992 15l-2.007-3l-4.015 8c-2.212-3.061-2.625-7.098-.915-10.463A10.14 10.14 0 0 1 12 4a10.14 10.14 0 0 1 8.945 5.537c1.71 3.365 1.297 7.402-.915 10.463l-4.517-8l-1.505 1.5"/><path d="m14.5 15l-3-6L9 13.5"/></g>`),
		g.Group(children),
	)
}