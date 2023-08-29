package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiffModifiedOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.2 3H4.8A1.8 1.8 0 0 0 3 4.8v14.4A1.8 1.8 0 0 0 4.8 21h14.4a1.8 1.8 0 0 0 1.8-1.8V4.8A1.8 1.8 0 0 0 19.2 3ZM19 19H5V5h14Z"/><circle cx="12" cy="12" r="4" fill="currentColor"/>`),
		g.Group(children),
	)
}