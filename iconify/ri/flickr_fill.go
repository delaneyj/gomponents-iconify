package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlickrFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 12a5 5 0 1 1-10 0a5 5 0 0 1 10 0Zm12 0a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/>`),
		g.Group(children),
	)
}