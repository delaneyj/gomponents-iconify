package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationArrowRightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 7c0 2.876-3.1 6.01-4.844 7.544a1.736 1.736 0 0 1-2.312 0C5.101 13.01 2 9.876 2 7a6 6 0 1 1 12 0ZM8.146 4.854L9.293 6H5.5a.5.5 0 0 0 0 1h3.793L8.146 8.146a.5.5 0 1 0 .708.708l2-2a.5.5 0 0 0 0-.708l-2-2a.5.5 0 1 0-.708.708Z"/>`),
		g.Group(children),
	)
}