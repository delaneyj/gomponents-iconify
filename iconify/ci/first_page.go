package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.444 18.01L10.432 12l6.012-6.01l1.415 1.414l-4.6 4.6l4.6 4.6l-1.415 1.406Zm-8.3-.01h-2V6h2v12Z"/>`),
		g.Group(children),
	)
}