package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWrapSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 3.5a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 0 1h-11a.5.5 0 0 1-.5-.5Zm0 4a.5.5 0 0 1 .5-.5h10a2.5 2.5 0 0 1 0 5H9.707l.647.646a.5.5 0 0 1-.708.708l-1.5-1.5a.5.5 0 0 1 0-.708l1.5-1.5a.5.5 0 0 1 .708.708L9.707 11H12.5a1.5 1.5 0 0 0 0-3h-10a.5.5 0 0 1-.5-.5ZM6 11a.5.5 0 0 1 0 1H2.5a.5.5 0 0 1 0-1H6Z"/>`),
		g.Group(children),
	)
}