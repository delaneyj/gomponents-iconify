package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Code(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.424 2.023a.75.75 0 0 1 .556.904L7.48 13.42a.75.75 0 0 1-1.46-.348L8.52 2.58a.75.75 0 0 1 .904-.556ZM11.16 4.22a.75.75 0 0 1 1.06 0l3.25 3.25L16 8l-.53.53l-3.25 3.25a.75.75 0 1 1-1.06-1.06L13.88 8l-2.72-2.72a.75.75 0 0 1 0-1.06ZM4.84 5.28a.75.75 0 1 0-1.06-1.06L.53 7.47L0 8l.53.53l3.25 3.25a.75.75 0 0 0 1.06-1.06L2.12 8l2.72-2.72Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}