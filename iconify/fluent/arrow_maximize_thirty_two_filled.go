package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMaximizeThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 4.25c0-.69.56-1.25 1.25-1.25h9.5c.69 0 1.25.56 1.25 1.25v9.5a1.25 1.25 0 1 1-2.5 0V7.268L7.268 26.5h6.482a1.25 1.25 0 1 1 0 2.5h-9.5C3.56 29 3 28.44 3 27.75v-9.5a1.25 1.25 0 1 1 2.5 0v6.482L24.732 5.5H18.25c-.69 0-1.25-.56-1.25-1.25Z"/>`),
		g.Group(children),
	)
}