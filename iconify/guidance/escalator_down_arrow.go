package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorDownArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m23.5 10.5l-10-10m10 10c-.398-.398-.654-1.133-.813-1.79a7.746 7.746 0 0 1-.148-2.687c.08-.697.235-1.464.544-1.773m.417 6.25c-.398-.398-1.133-.654-1.79-.812a7.745 7.745 0 0 0-2.687-.149c-.697.08-1.464.235-1.773.544M18.485 23.5H23.5v-6H18L9.757 9.257A6 6 0 0 0 5.515 7.5H.5v6H6l8.243 8.243a6 6 0 0 0 4.242 1.757Z"/>`),
		g.Group(children),
	)
}