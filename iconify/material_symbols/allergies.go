package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Allergies(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm-4.9-8.55l1.8-.9l-2-4l-1.8.9l2 4ZM9.5 18h2v-7.225L9.4 6.55l-1.8.9l1.9 3.8V18Zm3 0h2v-6.775l1.9-3.775l-1.8-.9l-2.1 4.2V18Zm4.4-4.55l2-4l-1.8-.9l-2 4l1.8.9Z"/>`),
		g.Group(children),
	)
}