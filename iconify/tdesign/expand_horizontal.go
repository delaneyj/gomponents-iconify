package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 4v16h-2V4h2Zm.586 8L9 7.086L10.414 8.5l-2.5 2.5h8.172l-2.5-2.5L15 7.086L19.914 12L15 16.914L13.586 15.5l2.5-2.5H7.914l2.5 2.5L9 16.914L4.086 12ZM22.5 4v16h-2V4h2Z"/>`),
		g.Group(children),
	)
}