package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersK(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20v-6.057l5 3.572V20h2v-2.743a1.5 1.5 0 0 0-.628-1.22L10.72 12l5.652-4.037A1.5 1.5 0 0 0 17 6.743V4h-2v2.485l-5 3.572V4H8v16h2Z"/>`),
		g.Group(children),
	)
}