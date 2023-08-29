package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.558 18.01l-1.414-1.413l4.6-4.6l-4.6-4.593L7.558 5.99L13.569 12l-6.01 6.01h-.001Zm10.3-.01h-2V6h2v12Z"/>`),
		g.Group(children),
	)
}