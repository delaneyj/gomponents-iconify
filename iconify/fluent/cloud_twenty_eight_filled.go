package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 6a6.5 6.5 0 0 0-6.481 6H7.25a4.75 4.75 0 1 0 0 9.5h13.5a4.75 4.75 0 1 0 0-9.5h-.269A6.5 6.5 0 0 0 14 6Z"/>`),
		g.Group(children),
	)
}