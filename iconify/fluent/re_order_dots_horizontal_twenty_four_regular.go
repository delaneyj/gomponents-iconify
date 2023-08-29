package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReOrderDotsHorizontalTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 15.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0-7a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm7 7a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0-7a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm7 7a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0-7a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/>`),
		g.Group(children),
	)
}