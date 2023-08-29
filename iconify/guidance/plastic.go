package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plastic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10 .5h4M6.5 10H7l1.027.171a24.168 24.168 0 0 0 7.946 0L17 10h.5m-11 7H7l1.027-.171a24.168 24.168 0 0 1 7.946 0L17 17h.5m-4-13.5h-3a7.236 7.236 0 0 0-4 6.472V21.5a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V9.972a7.236 7.236 0 0 0-4-6.472Z"/>`),
		g.Group(children),
	)
}