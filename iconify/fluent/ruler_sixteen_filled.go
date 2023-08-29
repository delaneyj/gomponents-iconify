package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 2A1.5 1.5 0 0 1 11 3.5v9A1.5 1.5 0 0 1 9.5 14h-3A1.5 1.5 0 0 1 5 12.5v-9A1.5 1.5 0 0 1 6.5 2h3Zm-2 9a.5.5 0 0 0 0-1H6v1h1.5ZM8 8.5a.5.5 0 0 0 0-1H6v1h2ZM7.5 6a.5.5 0 0 0 0-1H6v1h1.5Z"/>`),
		g.Group(children),
	)
}