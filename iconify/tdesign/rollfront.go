package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rollfront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.07 14A7.002 7.002 0 0 0 10 20h8.5v-2H10a5 5 0 0 1 0-10h7.086l-2.5 2.5L16 11.914L20.914 7L16 2.086L14.586 3.5l2.5 2.5H10a7 7 0 0 0-7 7v1h.07Z"/>`),
		g.Group(children),
	)
}