package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RocketTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.502 20.005A6.522 6.522 0 0 1 12 23.507a6.522 6.522 0 0 1-3.502-3.502h2.26c.326.488.747.912 1.242 1.243c.495-.33.916-.755 1.242-1.243h2.26ZM18 14.81l2 2.268v1.927H4v-1.927l2-2.268V9.005C6 5.52 8.504 2.558 12 1.46c3.496 1.098 6 4.061 6 7.545v5.804Zm-.73 2.195L16 15.565v-6.56c0-2.318-1.57-4.43-4-5.421c-2.43.99-4 3.103-4 5.42v6.561l-1.27 1.44h10.54Zm-5.27-6a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}