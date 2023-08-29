package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01M8 8h.01M8 16h.01M16 16h.01"/></g>`),
		g.Group(children),
	)
}