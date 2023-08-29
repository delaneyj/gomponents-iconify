package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialpadTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 6.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm6.5-12a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm6.5-18a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}