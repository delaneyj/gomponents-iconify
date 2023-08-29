package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtSymbol(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.243 5.757a6 6 0 1 0-.986 9.284a1 1 0 1 1 1.087 1.678A8 8 0 1 1 18 10a3 3 0 0 1-4.8 2.401A4 4 0 1 1 14 10a1 1 0 1 0 2 0c0-1.537-.586-3.07-1.757-4.243ZM12 10a2 2 0 1 0-4 0a2 2 0 0 0 4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}