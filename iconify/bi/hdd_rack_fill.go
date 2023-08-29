package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HddRackFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M2 2a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h1v2H2a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-1a2 2 0 0 0-2-2h-1V7h1a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H2zm.5 3a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm2 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm-2 7a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm2 0a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zM12 7v2H4V7h8z"/>`),
		g.Group(children),
	)
}