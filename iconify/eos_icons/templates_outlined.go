package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemplatesOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 0H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2Zm0 22H4V2h16Z"/><path fill="currentColor" d="M6 4h12v2H6zm0 4h7v2H6zm2 12h8l-4-3l-4 3zm11-1v-6l-5 3l5 3zM6 13v6l4-3l-4-3zm10-1H8l4 3l4-3z"/>`),
		g.Group(children),
	)
}