package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 3.5H4v-2h16v2Zm-8 .586L16.914 9L15.5 10.414l-2.5-2.5v8.172l2.5-2.5L16.914 15L12 19.914L7.086 15L8.5 13.586l2.5 2.5V7.914l-2.5 2.5L7.086 9L12 4.086ZM20 22.5H4v-2h16v2Z"/>`),
		g.Group(children),
	)
}