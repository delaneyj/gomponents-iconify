package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pyramid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 2.96l4.44 7.893l1.06-1.928L23.69 22H.29L11 2.96ZM13.227 11L11 7.04L8.772 11h4.455Zm-5.58 2L3.71 20H10v-7H7.647ZM12 13v4.107L14.259 13h-2.26Zm.691 7H15.5v-3h-1.16l-1.65 3Zm4.809 0h2.809l-1.65-3h-1.16v3Zm.059-5l-1.06-1.925L15.442 15h2.118Z"/>`),
		g.Group(children),
	)
}