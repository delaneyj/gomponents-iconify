package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Walking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M304 73c0-40-33-73-74-73c-40 0-72 33-72 73c0 41 32 73 72 73c41 0 74-32 74-73zm-79 644h-65s-3-185 2-207c4-22 47-68 49-86c2-19-17-97-17-97s-43 39-57 46s-124 25-124 25L0 342s105-23 120-35c15-11 55-108 85-120c20-9 74-5 101-5c28 0 120 53 129 65c10 11 65 119 65 119l-50 28l-40-86s-31-26-43-30c-13-4-30-9-33 2c-4 10 23 92 30 115c6 24 39 133 51 156s83 125 83 125l-61 40s-87-118-98-135c-12-17-46-108-46-108s-58 71-63 93c-4 20-5 151-5 151z"/>`),
		g.Group(children),
	)
}