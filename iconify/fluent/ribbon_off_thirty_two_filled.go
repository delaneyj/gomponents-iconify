package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RibbonOffThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.758 8.173A9.97 9.97 0 0 0 6 12a9.98 9.98 0 0 0 4 8a9.955 9.955 0 0 0 6 2c1.355 0 2.648-.27 3.827-.759l1.135 1.136A11.455 11.455 0 0 1 16 23.5c-2.199 0-4.253-.617-6-1.687V29a1 1 0 0 0 1.514.858L16 27.166l4.485 2.692A1 1 0 0 0 22 29v-5.586l6.293 6.293a1 1 0 0 0 1.414-1.414l-26-26a1 1 0 0 0-1.415 1.414l4.466 4.466ZM8.7 5.166L22.834 19.3A9.972 9.972 0 0 0 26 12c0-5.523-4.477-10-10-10a9.972 9.972 0 0 0-7.3 3.166Z"/>`),
		g.Group(children),
	)
}