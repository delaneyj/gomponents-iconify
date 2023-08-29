package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepBackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 5v22h2V16.5l.438.313l13 9L25 26.905V5.094l-1.563 1.093l-13 9L10 15.5V5zm15 3.906v14.188L12.75 16z"/>`),
		g.Group(children),
	)
}