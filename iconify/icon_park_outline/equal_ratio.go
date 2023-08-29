package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualRatio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="36" height="36" x="6" y="6" stroke="currentColor" stroke-linejoin="round" stroke-width="4" rx="3"/><path fill="currentColor" fill-rule="evenodd" d="M24 22.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0 8a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15.5 17v14m17-14v14"/></g>`),
		g.Group(children),
	)
}