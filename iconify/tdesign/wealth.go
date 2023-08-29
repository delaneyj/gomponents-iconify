package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wealth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.75 2.5h10.5v2.75c0 1.57-.69 2.98-1.783 3.942a9.034 9.034 0 0 1 4.41 3.951l.485.875l-1.75.97l-.484-.875A7 7 0 0 0 5 17.5v.5a2 2 0 0 0 2 2h7v2H7a4 4 0 0 1-4-4v-.5a9.003 9.003 0 0 1 5.533-8.308A5.237 5.237 0 0 1 6.75 5.25V2.5Zm5.25 6a3.25 3.25 0 0 0 3.25-3.25V4.5h-6.5v.75A3.25 3.25 0 0 0 12 8.5Zm4 7.5h7v2h-7v-2Zm0 4h7v2h-7v-2Z"/>`),
		g.Group(children),
	)
}