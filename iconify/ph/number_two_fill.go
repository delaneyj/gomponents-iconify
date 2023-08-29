package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 24H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16Zm-40 152a8 8 0 0 1 0 16H96a8 8 0 0 1-5.79-13.52L145.9 120a24 24 0 0 0-35.73-32a23.33 23.33 0 0 0-3.17 4.38a8 8 0 0 1-14-7.77a40.22 40.22 0 0 1 5.28-7.38a40 40 0 0 1 59.45 53.54l-.16.16L114.66 176Z"/>`),
		g.Group(children),
	)
}