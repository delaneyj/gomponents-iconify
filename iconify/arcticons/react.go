package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func React(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.44 24.76c21.734 0 11.732 18.335 23.027 18.335c3.876 0 5.24-2.934 5.24-2.934M11.44 4.5v39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.293 4.5h16.011c6.054 0 10.96 4.535 10.96 10.13c0 0 0 0 0 0c0 5.594-4.906 10.13-10.96 10.13H11.46M8.293 43.5h6.293"/>`),
		g.Group(children),
	)
}