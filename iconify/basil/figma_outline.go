package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.75 15.373a3.999 3.999 0 0 0 6.195-4.654A3.999 3.999 0 0 0 17.582 9a3.996 3.996 0 0 0 1.668-3.25a4 4 0 0 0-4-4h-6.5A4 4 0 0 0 6.418 9a4 4 0 0 0 0 6.5a4 4 0 1 0 6.332 3.25v-3.377Zm-4-12.123a2.5 2.5 0 1 0 0 5h2.5v-5h-2.5Zm2.5 13h-2.5a2.5 2.5 0 1 0 2.5 2.5v-2.5Zm-2.5-6.5a2.5 2.5 0 0 0 0 5h2.5v-5h-2.5Zm4 2.5a2.5 2.5 0 1 0 5.001 0a2.5 2.5 0 0 0-5.001 0Zm2.5-4a2.5 2.5 0 1 0 0-5h-2.5v5h2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}