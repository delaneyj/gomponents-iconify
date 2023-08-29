package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CelebrationOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22L7 8l9 9l-14 5Zm3.3-3.3l7.05-2.5l-4.55-4.55l-2.5 7.05Zm9.25-6.15L13.5 11.5l7.525-7.525L23.55 6.5L22.5 7.55l-1.475-1.475l-6.475 6.475Zm-4-4L9.5 7.5l1.45-1.45l-1.5-1.5L10.5 3.5l2.55 2.55l-2.5 2.5Zm2 2L11.5 9.5l4.475-4.475L13.5 2.55l1.05-1.05l3.525 3.525l-5.525 5.525Zm4 4L15.5 13.5l3.525-3.525L22.55 13.5l-1.05 1.05l-2.475-2.475l-2.475 2.475ZM5.3 18.7Z"/>`),
		g.Group(children),
	)
}