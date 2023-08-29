package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dcr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#2ED6A1"/><path fill="#FFF" fill-rule="nonzero" d="M14.856 17.62h4.42a2.987 2.987 0 0 0 2.978-2.996a2.987 2.987 0 0 0-2.978-2.995H17.87L14.856 9h4.42a5.602 5.602 0 0 1 5.488 4.545a5.632 5.632 0 0 1-3.379 6.29l3.485 3.042h-3.991l-6.023-5.258zm2.157-3.36h-4.42a2.987 2.987 0 0 0-2.978 2.994a2.987 2.987 0 0 0 2.978 2.995H14l3.012 2.63h-4.419a5.602 5.602 0 0 1-5.488-4.545a5.632 5.632 0 0 1 3.379-6.29L7.002 9h3.991l6.02 5.26z"/></g>`),
		g.Group(children),
	)
}