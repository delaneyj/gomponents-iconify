package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsEditLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 2C7.589 2 4 5.589 4 9.995C3.971 16.44 11.696 21.784 12 22c0 0 8.029-5.56 8-12c0-4.411-3.589-8-8-8zM9.799 14.987H8v-1.799l4.977-4.97l1.799 1.799l-4.977 4.97zm5.824-5.817l-1.799-1.799L15.196 6l1.799 1.799l-1.372 1.371z" fill="currentColor"/>`),
		g.Group(children),
	)
}