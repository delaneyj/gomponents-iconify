package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftRightThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.914 7.5L15.5 12.914L14.086 11.5l3-3H8.5v-2h8.586l-3-3L15.5 2.086L20.914 7.5Zm-5.414 10H6.914l3 3L8.5 21.914L3.086 16.5L8.5 11.086L9.914 12.5l-3 3H15.5v2Z"/>`),
		g.Group(children),
	)
}