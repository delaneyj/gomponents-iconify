package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OvalThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 6C6.477 6 2 10.477 2 16s4.477 10 10 10h8c5.523 0 10-4.477 10-10S25.523 6 20 6h-8Z"/>`),
		g.Group(children),
	)
}