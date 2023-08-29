package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RhombusThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.524 5a3.25 3.25 0 0 0-3.018 2.043l-6.2 15.5C.452 24.678 2.024 27 4.324 27h17.153a3.25 3.25 0 0 0 3.017-2.043l6.2-15.5C31.548 7.322 29.976 5 27.677 5H10.524Z"/>`),
		g.Group(children),
	)
}