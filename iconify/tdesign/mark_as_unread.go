package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkAsUnread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.5 1.34l7.744 4.555L16.23 7.62L9.5 3.66L2 8.072V20H0V6.928L9.5 1.34ZM4 9h19v14H4V9Zm3.992 2l5.508 3.787L19.008 11H7.992ZM21 12.057l-7.5 5.157L6 12.057V21h15v-8.943Z"/>`),
		g.Group(children),
	)
}