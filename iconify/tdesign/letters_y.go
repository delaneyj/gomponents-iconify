package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersY(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.713 8.22A1.5 1.5 0 0 0 17 7.337V4h-2v3.175L12 11.3L9 7.175V4H7v3.337c0 .317.1.626.287.883L11 13.325V20h2v-6.675l3.713-5.105Z"/>`),
		g.Group(children),
	)
}