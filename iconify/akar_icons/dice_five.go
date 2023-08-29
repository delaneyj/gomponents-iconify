package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" stroke-linejoin="round" rx="2"/><path d="M8 8h.5m7 0h.5m-4 4h.5m-.25-.25v.5M8 16h.5m7 0h.5M8.25 7.75v.5m0 7.5v.5m7.5-8.5v.5m0 7.5v.5"/></g>`),
		g.Group(children),
	)
}