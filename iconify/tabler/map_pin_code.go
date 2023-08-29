package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPinCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 11a3 3 0 1 0 6 0a3 3 0 0 0-6 0"/><path d="M11.85 21.48a1.992 1.992 0 0 1-1.263-.58l-4.244-4.243a8 8 0 1 1 13.385-3.585M20 21l2-2l-2-2m-3 0l-2 2l2 2"/></g>`),
		g.Group(children),
	)
}