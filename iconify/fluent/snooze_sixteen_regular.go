package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnoozeSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.001 8h2.5a.5.5 0 0 1 .432.753l-.048.067L5.068 11h1.433a.5.5 0 0 1 .09.992L6.5 12H4a.5.5 0 0 1-.432-.753l.048-.067L5.433 9H4.001a.5.5 0 0 1-.09-.992L4.001 8h2.5h-2.5Zm5-5h3.5a.5.5 0 0 1 .452.714l-.042.073L9.96 8h2.54a.5.5 0 0 1 .09.992L12.5 9H9a.5.5 0 0 1-.452-.714l.042-.073L11.541 4H9a.5.5 0 0 1-.09-.992L9.001 3h3.5h-3.5Z"/>`),
		g.Group(children),
	)
}