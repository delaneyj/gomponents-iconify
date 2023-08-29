package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M22 6a4 4 0 0 1 8 0v1a1 1 0 1 0 2 0V6a6 6 0 0 0-12 0v3H8.5A4.5 4.5 0 0 0 4 13.5v12A4.5 4.5 0 0 0 8.5 30h15a4.5 4.5 0 0 0 4.5-4.5v-12A4.5 4.5 0 0 0 23.5 9H22V6ZM8.5 11h15a2.5 2.5 0 0 1 2.5 2.5v12a2.5 2.5 0 0 1-2.5 2.5h-15A2.5 2.5 0 0 1 6 25.5v-12A2.5 2.5 0 0 1 8.5 11Zm9.5 9a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`),
		g.Group(children),
	)
}