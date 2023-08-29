package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClosedThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 8a6 6 0 0 1 12 0v1h1.5a4.5 4.5 0 0 1 4.5 4.5v12a4.5 4.5 0 0 1-4.5 4.5h-15A4.5 4.5 0 0 1 4 25.5v-12A4.5 4.5 0 0 1 8.5 9H10V8Zm6-4a4 4 0 0 0-4 4v1h8V8a4 4 0 0 0-4-4Zm0 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}