package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareRootAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m19.313 6l-.25.656l-6.157 16.063l-3-6.157L9.625 16H6v2h2.375L12 25.438l.281.562h1.5l.25-.656L20.687 8H26V6z"/>`),
		g.Group(children),
	)
}