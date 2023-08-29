package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrowUpSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 7c0 2.876-3.1 6.01-4.844 7.544a1.736 1.736 0 0 1-2.312 0C5.101 13.01 2 9.876 2 7a6 6 0 1 1 12 0Zm-7.646-.146L7.5 5.707V9.5a.5.5 0 0 0 1 0V5.707l1.146 1.147a.5.5 0 0 0 .708-.708l-2-2a.5.5 0 0 0-.708 0l-2 2a.5.5 0 1 0 .708.708Z"/>`),
		g.Group(children),
	)
}