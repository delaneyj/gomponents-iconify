package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m16.75 8.412l24.417 12.705a3.25 3.25 0 0 1 0 5.766L16.75 39.588A3.25 3.25 0 0 1 12 36.705v-25.41a3.25 3.25 0 0 1 4.549-2.98l.201.097Zm-1.154 2.218l-.11-.047a.75.75 0 0 0-.986.712v25.41a.749.749 0 0 0 1.096.665l24.417-12.705a.75.75 0 0 0 0-1.33L15.596 10.63Z"/>`),
		g.Group(children),
	)
}