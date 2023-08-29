package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.5 10a5.5 5.5 0 1 0 11 0a5.5 5.5 0 0 0-11 0Zm9 0a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}