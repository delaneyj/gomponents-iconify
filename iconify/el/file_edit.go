package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M284.506 0L58.436 218.249V1200h850.418V817.511L803.861 922.504v173.167H163.428V292.179h197.59v-187.85h442.844v272.996L490.433 690.753L397.32 977.619l286.793-93.188l457.452-457.452L947.885 233.3l-39.031 39.031V0H284.506zm254.402 739.154l96.803 96.876l-143.434 46.557l46.631-143.433z"/>`),
		g.Group(children),
	)
}