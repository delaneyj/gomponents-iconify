package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jump(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="m8 10l12 10.254v9.714L10.857 44M40 10L28 20.254v9.714L37.143 44"/><circle cx="24" cy="8" r="4"/></g>`),
		g.Group(children),
	)
}