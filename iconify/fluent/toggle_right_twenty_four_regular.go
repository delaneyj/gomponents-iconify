package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleRightTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.75 14.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM2 12a5 5 0 0 1 5-5h10a5 5 0 0 1 0 10H7a5 5 0 0 1-5-5Zm5-3.5a3.5 3.5 0 1 0 0 7h10a3.5 3.5 0 1 0 0-7H7Z"/>`),
		g.Group(children),
	)
}