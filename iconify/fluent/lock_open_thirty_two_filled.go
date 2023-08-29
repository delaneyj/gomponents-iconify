package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22 6a4 4 0 0 1 8 0v1a1 1 0 1 0 2 0V6a6 6 0 0 0-12 0v3H8.5A4.5 4.5 0 0 0 4 13.5v12A4.5 4.5 0 0 0 8.5 30h15a4.5 4.5 0 0 0 4.5-4.5v-12A4.5 4.5 0 0 0 23.5 9H22V6Zm-6 16a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}