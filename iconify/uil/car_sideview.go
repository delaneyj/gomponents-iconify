package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarSideview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 9.5h-.32l-1.25-3.12a3 3 0 0 0-2.78-1.88h-6A3 3 0 0 0 5.7 6.91L5.18 9.5H5a3 3 0 0 0-3 3v3a1 1 0 0 0 1 1h1a3 3 0 0 0 6 0h4a3 3 0 0 0 6 0h1a1 1 0 0 0 1-1v-3a3 3 0 0 0-3-3Zm-6-3h1.65a1 1 0 0 1 .92.63l.95 2.37H13Zm-5.34.8a1 1 0 0 1 1-.8H11v3H7.22ZM7 17.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm10 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm3-3h-.78a3 3 0 0 0-4.44 0H9.22a3 3 0 0 0-4.44 0H4v-2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}