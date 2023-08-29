package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewPicture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" rx="3"/><path stroke-linecap="round" d="m6 28l10.693-9.802a2 2 0 0 1 2.653-.044L32 29m-4-5l4.773-4.773a2 2 0 0 1 2.615-.186L42 24M6 20v10m36-10v10"/></g>`),
		g.Group(children),
	)
}