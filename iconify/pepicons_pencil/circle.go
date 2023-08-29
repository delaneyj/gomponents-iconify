package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 5a5 5 0 1 0 0 10a5 5 0 0 0 0-10Zm-6 5a6 6 0 1 1 12 0a6 6 0 0 1-12 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}