package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DlfNachrichten(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.715 19.586h14.682M19.668 25.68H39.61v-2.3H19.668v2.298m22.48 3.797v4.135H19.72v-4.135ZM16.67 22.77a9.197 9.197 0 1 1 7.002-5.803M19.715 43.5h14.682m-14.729-6.093H39.61v2.298H19.668v-2.298"/>`),
		g.Group(children),
	)
}