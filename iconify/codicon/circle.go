package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8zm2.61-4a2.61 2.61 0 1 1-5.22 0a2.61 2.61 0 0 1 5.22 0zM8 5.246z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}