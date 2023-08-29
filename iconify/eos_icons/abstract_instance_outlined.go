package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbstractInstanceOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21.732 11l-4-7a2 2 0 0 0-1.74-1h-8a2 2 0 0 0-1.71 1l-4 7a2 2 0 0 0 0 2l4 7a2 2 0 0 0 1.71 1h8a2 2 0 0 0 1.74-1l4-7a2 2 0 0 0 0-2ZM16 19H8l-4-7l4-7h8l4 7Z"/><circle cx="12" cy="12" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}