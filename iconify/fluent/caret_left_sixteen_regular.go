package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretLeftSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.429 11.864a1 1 0 0 0 1.57-.821V4.957a1 1 0 0 0-1.57-.821L5.644 6.769a1.5 1.5 0 0 0 0 2.462l3.785 2.633Zm.57-.821L6.216 8.41a.5.5 0 0 1 0-.82L10 4.956v6.086Z"/>`),
		g.Group(children),
	)
}