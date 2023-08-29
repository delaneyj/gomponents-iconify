package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipPrevious(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 18l-8.5-6L18 6v12ZM8 18H6V6h2v12Z"/>`),
		g.Group(children),
	)
}