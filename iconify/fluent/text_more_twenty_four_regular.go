package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextMoreTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.75 2a.75.75 0 0 1 .69.457l6.972 16.431a2.495 2.495 0 0 0-2.006-.886L13.92 14.5H5.58l-2.14 5.043a.75.75 0 0 1-1.38-.586l7-16.5A.75.75 0 0 1 9.75 2Zm3.534 11L9.75 4.67L6.216 13h7.068ZM12 20.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm3.5 1.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm5 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}