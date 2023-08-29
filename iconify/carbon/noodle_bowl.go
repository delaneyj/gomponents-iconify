package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoodleBowl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m11.414 15l-8-8L2 8.414L8.586 15H2v1a14 14 0 0 0 28 0v-1ZM16 28A12.017 12.017 0 0 1 4.041 17H27.96A12.017 12.017 0 0 1 16 28Z"/><path fill="currentColor" d="M22 8a5.005 5.005 0 0 0-1.57.255A8.024 8.024 0 0 0 14 5a7.936 7.936 0 0 0-4.906 1.68L4.414 2L3 3.414l6.05 6.05l.707-.707A5.96 5.96 0 0 1 14 7a6.02 6.02 0 0 1 4.688 2.264a5.06 5.06 0 0 0-.59.61A2.99 2.99 0 0 1 15.754 11H12v2h3.754a4.98 4.98 0 0 0 3.904-1.875A3 3 0 0 1 25 13h2a5.006 5.006 0 0 0-5-5Z"/>`),
		g.Group(children),
	)
}