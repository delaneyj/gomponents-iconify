package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Delta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2.091V22.79L3.172 13.05L19 2.091ZM6.828 12.951L17 19.21V5.907L6.828 12.951Z"/>`),
		g.Group(children),
	)
}