package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forbidden(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5zM1.25 8A6 6 0 0 1 3 3.85L12.09 13A7.12 7.12 0 0 1 8 14.25A6.52 6.52 0 0 1 1.25 8zM13 12.15L3.91 3A7.12 7.12 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6 6 0 0 1 13 12.15z"/>`),
		g.Group(children),
	)
}