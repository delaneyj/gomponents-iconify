package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUpTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.161 14.907c-.707.809-.133 2.073.941 2.073h11.796c1.074 0 1.648-1.264.941-2.073l-5.522-6.31a1.75 1.75 0 0 0-2.634 0l-5.522 6.31Zm1.492.573l5.159-5.896a.25.25 0 0 1 .376 0l5.16 5.896H6.652Z"/>`),
		g.Group(children),
	)
}