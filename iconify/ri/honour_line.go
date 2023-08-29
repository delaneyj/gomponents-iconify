package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HonourLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 4v14.721a.5.5 0 0 1-.298.458L12 23.03l-8.702-3.85A.5.5 0 0 1 3 18.722V4H1V2h22v2h-2ZM5 4v13.745l7 3.1l7-3.1V4H5Zm3 4h8v2H8V8Zm0 4h8v2H8v-2Z"/>`),
		g.Group(children),
	)
}