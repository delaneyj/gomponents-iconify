package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraDomeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 3.5A1.5 1.5 0 0 1 2.5 2h11a1.5 1.5 0 0 1 0 3h-11A1.5 1.5 0 0 1 1 3.5ZM8 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm-1 2a1 1 0 1 1 2 0a1 1 0 0 1-2 0Zm7-4H2v3a6 6 0 0 0 12 0V6Zm-9 4a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/>`),
		g.Group(children),
	)
}