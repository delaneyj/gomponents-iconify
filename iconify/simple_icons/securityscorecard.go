package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Securityscorecard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.37 2.5L12 5L7.63 7.5v-5L12 0ZM22.487 6l.003 4.986l-8.728-4.993L18.118 3.5Zm-4.369 9.508l-.001 4.997l4.377-2.5l-.003-5.018l-4.373-2.502zm-10.49 5.994L12 24l4.369-2.495v-4.997zM7.63 9.5v5.001l4.37 2.5l4.37-2.494V9.5L12 7Zm-6.124 8.504L1.508 13l8.747 5.003l-4.376 2.5ZM5.882 3.5l-4.37 2.501l-.002 4.998l4.372 2.503z"/>`),
		g.Group(children),
	)
}