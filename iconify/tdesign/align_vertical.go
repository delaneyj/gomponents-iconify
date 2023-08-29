package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 6.586V0h-2v6.586L9.707 5.293L9 4.586L7.586 6l.707.707l3 3l.707.707l.707-.707l3-3L16.414 6L15 4.586l-.707.707L13 6.586ZM21 13H3v-2h18v2Zm-9 .586l.707.707l3 3l.707.707L15 19.414l-.707-.707L13 17.414V24h-2v-6.586l-1.293 1.293l-.707.707L7.586 18l.707-.707l3-3l.707-.707Z"/>`),
		g.Group(children),
	)
}