package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShrinkHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.5 11h5.086l-2.5-2.5L5.5 7.086L10.414 12L5.5 16.914L4.086 15.5l2.5-2.5H1.5v-2ZM13 3v18h-2V3h2Zm.586 9L18.5 7.086L19.914 8.5l-2.5 2.5H22.5v2h-5.086l2.5 2.5l-1.414 1.414L13.586 12Z"/>`),
		g.Group(children),
	)
}