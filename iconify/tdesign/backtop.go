package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backtop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4h16v2H4V4Zm8 3.586l6.914 6.914l-1.414 1.414l-4.5-4.5V21h-2v-9.586l-4.5 4.5L5.086 14.5L12 7.586Z"/>`),
		g.Group(children),
	)
}