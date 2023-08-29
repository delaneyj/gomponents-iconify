package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TipsDouble(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7H6v11h4.914l2.586 2.586L16.086 18H21V7Zm2 13h-6.086L13.5 23.414L10.086 20H4V5h19v15ZM19 3.5H2.5v12h-2v-14H19v2Z"/>`),
		g.Group(children),
	)
}