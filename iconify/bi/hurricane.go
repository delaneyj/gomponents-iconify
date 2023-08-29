package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hurricane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6.999 2.6A5.5 5.5 0 0 1 15 7.5a.5.5 0 0 0 1 0a6.5 6.5 0 1 0-13 0a5 5 0 0 0 6.001 4.9A5.5 5.5 0 0 1 1 7.5a.5.5 0 0 0-1 0a6.5 6.5 0 1 0 13 0a5 5 0 0 0-6.001-4.9zM10 7.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0z"/>`),
		g.Group(children),
	)
}