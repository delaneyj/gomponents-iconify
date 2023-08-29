package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotonePersonRemoveAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 16c2.69 0 5.77 1.28 6 2H4c.2-.71 3.3-2 6-2z" opacity=".3"/><circle cx="10" cy="8" r="2" fill="currentColor" opacity=".3"/><path fill="currentColor" d="M14 8c0-2.21-1.79-4-4-4S6 5.79 6 8s1.79 4 4 4s4-1.79 4-4zm-2 0c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2zM2 18v2h16v-2c0-2.66-5.33-4-8-4s-8 1.34-8 4zm2 0c.2-.71 3.3-2 6-2c2.69 0 5.77 1.28 6 2H4zm13-8h6v2h-6z"/>`),
		g.Group(children),
	)
}