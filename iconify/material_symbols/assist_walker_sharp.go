package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssistWalkerSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 7.5q-.825 0-1.413-.588T10.5 5.5q0-.825.588-1.413T12.5 3.5q.825 0 1.413.588T14.5 5.5q0 .825-.588 1.413T12.5 7.5ZM8 21v-5.25l-.775-.7l.175 1.35l-3.675 4.725L2.15 19.9l3.15-4.05l-2.075-4.1l4.2-4.15q.3-.3.662-.45T8.825 7q.6 0 .95.225t.475.35l2 1.975q.675.675 1.65 1.062T16 11h2.975l.8 7.7q.325.2.525.537t.2.763q0 .625-.438 1.063T19 21.5q-.625 0-1.075-.438T17.475 20q0-.425.2-.763t.55-.537l-.125-1.2h-3.25L14.5 21H13l.85-8.275q-1.425-.375-2.15-.988T10.25 10.4l-2.375 2.35L10 14.875V21H8Zm7-5h2.95l-.35-3.5h-2.225L15 16Z"/>`),
		g.Group(children),
	)
}