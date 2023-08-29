package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="3" y="10" rx="2"/><path d="M6 6a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v4H6V6Z"/></g>`),
		g.Group(children),
	)
}