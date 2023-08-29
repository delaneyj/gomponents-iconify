package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleHalfFillTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.543 9.25h12.914a6.501 6.501 0 0 0-12.914 0ZM2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0Z"/>`),
		g.Group(children),
	)
}