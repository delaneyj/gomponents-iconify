package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepForward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19.75a.75.75 0 0 1-.29-.06a.74.74 0 0 1-.46-.69V5a.74.74 0 0 1 .46-.69a.75.75 0 0 1 .82.16l7 7a.75.75 0 0 1 0 1.06l-7 7a.77.77 0 0 1-.53.22Zm.75-12.94v10.38L13.94 12Z"/><path fill="currentColor" d="M16 19.75a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v14a.76.76 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}