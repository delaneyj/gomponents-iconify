package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtmCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.666 3.07H2.333C1.597 3.07 1 3.641 1 4.347v.639h16v-.639c0-.705-.598-1.277-1.334-1.277zM1 11.648c0 .731.597 1.323 1.333 1.323h13.333c.736 0 1.334-.592 1.334-1.323V7.07H1v4.578zm6-2.745h4v1.125H7V8.903zm-2.042 2.03H13v1.096H4.958v-1.096z"/>`),
		g.Group(children),
	)
}