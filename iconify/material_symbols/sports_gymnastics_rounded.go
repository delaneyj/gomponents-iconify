package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsGymnasticsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 8q-.825 0-1.413-.588T4 6q0-.825.588-1.413T6 4q.825 0 1.413.588T8 6q0 .825-.588 1.413T6 8Zm6.95 14q-.4 0-.688-.275t-.312-.675L11.5 12L8 11H2q-.425 0-.713-.288T1 10q0-.425.288-.713T2 9h5l6.25-4.475q.325-.225.7-.175t.65.35q.275.35.225.762t-.4.688L11.15 8.5H14l7.15-4.125q.275-.175.613-.1t.612.4q.275.3.225.688T22.225 6L14.5 12l-.45 9.05q-.025.4-.325.675T12.95 22Z"/>`),
		g.Group(children),
	)
}