package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseCircleTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0Zm7-2.5a.5.5 0 0 0-1 0v5a.5.5 0 0 0 1 0v-5Zm3 0a.5.5 0 0 0-1 0v5a.5.5 0 0 0 1 0v-5Z"/>`),
		g.Group(children),
	)
}