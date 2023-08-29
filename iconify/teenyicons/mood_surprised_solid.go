package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoodSurprisedSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2H7Z"/><path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0ZM4 6h1V5H4v1Zm6 0h1V5h-1v1ZM5 9a2 2 0 0 1 2-2h1a2 2 0 1 1 0 4H7a2 2 0 0 1-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}