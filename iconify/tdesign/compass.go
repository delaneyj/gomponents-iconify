package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 12a9 9 0 1 0-18 0a9 9 0 0 0 18 0ZM12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1Zm5.435 5.565l-2.64 8.23l-8.23 2.64l2.64-8.23l8.23-2.64Zm-6.64 4.23l-1.138 3.548l3.547-1.138l1.138-3.548l-3.547 1.139Z"/>`),
		g.Group(children),
	)
}