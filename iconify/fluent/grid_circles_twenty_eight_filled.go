package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridCirclesTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 13A5 5 0 1 0 8 3a5 5 0 0 0 0 10Zm12 0a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm-7 7a5 5 0 1 1-10 0a5 5 0 0 1 10 0Zm7 5a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/>`),
		g.Group(children),
	)
}