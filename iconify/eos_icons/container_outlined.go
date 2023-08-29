package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContainerOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 10h-2l-4.79-3.5A1.854 1.854 0 0 0 14 5a2.03 2.03 0 0 0-1-1.721V2h-1v2a1 1 0 1 1-1 1h-1a1.846 1.846 0 0 0 .796 1.5L6 10H4a2.002 2.002 0 0 0-2 2v8a2.002 2.002 0 0 0 2 2h16a1.997 1.997 0 0 0 2-2v-8a2.002 2.002 0 0 0-2-2Zm-8-3c.006 0 4 3 4 3H8Zm8 13H4v-8h16Z"/><path fill="currentColor" d="M14 19a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1Zm4 0a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1ZM6 19a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1Zm4 0a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1Z"/>`),
		g.Group(children),
	)
}