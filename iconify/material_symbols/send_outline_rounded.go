package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.8 12.925l-15.4 6.5q-.5.2-.95-.088T3 18.5v-13q0-.55.45-.838t.95-.087l15.4 6.5q.625.275.625.925t-.625.925ZM5 17l11.85-5L5 7v3.5l6 1.5l-6 1.5V17Zm0 0V7v10Z"/>`),
		g.Group(children),
	)
}