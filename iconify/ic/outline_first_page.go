package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineFirstPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.41 16.59L13.82 12l4.59-4.59L17 6l-6 6l6 6l1.41-1.41zM6 6h2v12H6V6z"/>`),
		g.Group(children),
	)
}