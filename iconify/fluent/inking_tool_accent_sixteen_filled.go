package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InkingToolAccentSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.5 2v2a.5.5 0 0 1-.5.5H2a.5.5 0 0 1-.5-.5V2h13Zm-4 2.5l.52.476l-2.398 4.75a.5.5 0 1 1-.893-.451l2.17-4.3l.601-.475ZM9 13c0 .828-.448 1.5-1 1.5s-1-.672-1-1.5s.448-1.5 1-1.5s1 .672 1 1.5Z"/>`),
		g.Group(children),
	)
}