package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownLeftThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.75 29a1.25 1.25 0 1 0 0-2.5H7.268L28.634 5.134a1.25 1.25 0 0 0-1.768-1.768L5.5 24.732V15.25a1.25 1.25 0 0 0-2.5 0v12.5c0 .69.56 1.25 1.25 1.25h12.5Z"/>`),
		g.Group(children),
	)
}