package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 19.75a.77.77 0 0 1-.53-.22l-7-7a.75.75 0 0 1 0-1.06l7-7a.75.75 0 0 1 .82-.16a.74.74 0 0 1 .46.69v14a.74.74 0 0 1-.46.69a.75.75 0 0 1-.29.06ZM10.06 12l5.19 5.19V6.81Z"/><path fill="currentColor" d="M8 19.75a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v14a.76.76 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}