package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonWaxingCrescent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a9.91 9.91 0 0 0-3 .46a10 10 0 0 1 0 19.08A10 10 0 1 0 12 2Z"/>`),
		g.Group(children),
	)
}