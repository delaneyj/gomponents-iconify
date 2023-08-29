package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wheelchair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M29.622 35c-1.332 5.176-6.03 9-11.622 9c-6.627 0-12-5.373-12-12c0-4.843 2.869-9.016 7-10.912"/><path d="m18 12l2 18l15-1l3 11h3"/><path d="M22 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm3 12h8"/></g>`),
		g.Group(children),
	)
}