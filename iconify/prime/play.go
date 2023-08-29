package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 19.75a.75.75 0 0 1-.29-.06a.74.74 0 0 1-.46-.69V5A.75.75 0 0 1 9 4.47l7 7a.75.75 0 0 1 0 1.06l-7 7a.77.77 0 0 1-.5.22Zm.75-12.94v10.38L14.44 12Z"/>`),
		g.Group(children),
	)
}