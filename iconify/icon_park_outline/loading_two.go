package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 4v8m14.142-2.142l-5.657 5.657M44 24h-8m2.142 14.142l-5.657-5.657M24 44v-8M9.858 38.142l5.657-5.657M4 24h8M9.858 9.858l5.657 5.657m.832-9.993l1.53 3.696M5.522 16.346l3.696 1.53M5.522 31.654l3.696-1.531m7.129 12.355l1.53-3.696m13.777 3.696l-1.531-3.696m12.355-7.128l-3.696-1.531m3.696-13.777l-3.696 1.53M31.654 5.522l-1.531 3.696"/>`),
		g.Group(children),
	)
}