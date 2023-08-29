package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUpSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.136 9.427a1 1 0 0 0 .82 1.571h6.087a1 1 0 0 0 .82-1.571L9.232 5.643a1.5 1.5 0 0 0-2.462 0L4.136 9.427Zm.82.571L7.59 6.214a.5.5 0 0 1 .821 0l2.633 3.784H4.957Z"/>`),
		g.Group(children),
	)
}