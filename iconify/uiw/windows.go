package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20 10.873V20L8.479 18.537l.001-7.664H20Zm-13.12 0l-.001 7.461L0 17.461v-6.588h6.88ZM20 9.273H8.48l-.001-7.81L20 0v9.273ZM6.879 1.666l.001 7.607H0V2.539l6.879-.873Z"/>`),
		g.Group(children),
	)
}