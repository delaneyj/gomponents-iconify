package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 21q40-16 77-18q13 0 35 4q16 4 49 22l-42 142q-41-27-82-27q-35 0-77 16zm65 341q26-90 43-144q-9-6-17-10q-32-17-62-17q-1 0-4 .5t-4 .5q-35 3-63 14q-8 2-12 4L2 350q39-16 74-16q41 0 85 28zm246-104q-42 13-75 13q-55 0-93-28l-40 141q50 29 83 29q37 0 86-20zm55-191q-38 16-77 16q-51 0-91-28l-41 142q40 27 89 27q38 0 79-18v-1h1z"/>`),
		g.Group(children),
	)
}