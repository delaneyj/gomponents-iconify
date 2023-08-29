package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eurowings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.901 12.541c-3.353-.496-4.968 1.367-7.328 3.105L4.5 21.857l1.49.746c1.988.994 4.72.62 6.583-.745l11.054-8.695l-3.726-.62Zm13.787 4.348l-4.099-.746s-2.857-.745-5.465 1.118l-12.793 10.06l1.987 1.118c1.988.994 4.596.621 6.46-.745L33.811 16.89m9.689 4.097l-4.471-.745c-2.609-.497-3.975.373-6.459 2.36L18.784 33.78l1.987 1.118c1.987.994 4.782.87 6.583-.62L43.5 20.987Z"/>`),
		g.Group(children),
	)
}