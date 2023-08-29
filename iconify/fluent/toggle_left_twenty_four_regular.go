package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeftTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.25 14.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5ZM22 12a5 5 0 0 0-5-5H7a5 5 0 0 0 0 10h10a5 5 0 0 0 5-5Zm-5-3.5a3.5 3.5 0 1 1 0 7H7a3.5 3.5 0 1 1 0-7h10Z"/>`),
		g.Group(children),
	)
}