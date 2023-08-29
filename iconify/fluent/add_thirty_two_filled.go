package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.25 3c.69 0 1.25.56 1.25 1.25V15h10.75a1.25 1.25 0 1 1 0 2.5H17.5v10.75a1.25 1.25 0 1 1-2.5 0V17.5H4.25a1.25 1.25 0 1 1 0-2.5H15V4.25c0-.69.56-1.25 1.25-1.25Z"/>`),
		g.Group(children),
	)
}