package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.5 11.13l-14-8.08a1 1 0 0 0-1 0a1 1 0 0 0-.5.87v16.16a1 1 0 0 0 .5.87a1 1 0 0 0 1 0l14-8.08a1 1 0 0 0 0-1.74ZM6 18.35V5.65L17 12Z"/>`),
		g.Group(children),
	)
}