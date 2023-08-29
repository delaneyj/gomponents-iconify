package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongUpR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.793 4.61L12.068.398l4.21 4.275l-1.424 1.403l-1.804-1.831l-.07 11.886l3.227 3.226l-4.243 4.243l-4.242-4.243l3.259-3.26l.07-11.89l-1.854 1.826L7.793 4.61Zm4.171 16.163l1.414-1.415l-1.414-1.414l-1.414 1.414l1.414 1.415Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}