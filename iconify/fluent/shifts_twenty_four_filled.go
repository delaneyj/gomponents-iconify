package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShiftsTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.75 3A3.25 3.25 0 0 1 21 6.25v11.5A3.25 3.25 0 0 1 17.75 21H6.25A3.25 3.25 0 0 1 3 17.75V6.25A3.25 3.25 0 0 1 6.25 3h11.5Zm-6 3a.75.75 0 0 0-.75.75v6c0 .414.336.75.75.75h4.498a.75.75 0 0 0 0-1.5H12.5V6.75a.75.75 0 0 0-.75-.75Z"/>`),
		g.Group(children),
	)
}