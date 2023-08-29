package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SnoozeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 7.5h2.5a.75.75 0 0 1 .665 1.097l-.055.089L5.457 11h1.042a.75.75 0 0 1 .102 1.493L6.5 12.5H4a.75.75 0 0 1-.665-1.097l.055-.089L5.042 9H3.999a.75.75 0 0 1-.101-1.493l.101-.007h2.5H4Zm5-5h3.5a.75.75 0 0 1 .681 1.063l-.049.09L10.366 8h2.133a.75.75 0 0 1 .102 1.493L12.5 9.5H9a.75.75 0 0 1-.682-1.063l.05-.09L11.133 4H9a.75.75 0 0 1-.101-1.493L9 2.5h3.5H9Z"/>`),
		g.Group(children),
	)
}