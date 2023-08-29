package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 5.5a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9ZM3.5 10a6.5 6.5 0 1 1 13 0a6.5 6.5 0 0 1-13 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}