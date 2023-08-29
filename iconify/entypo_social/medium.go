package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.75 5S18 5 18 5.75v8.5s0 .75.75.75H19v2h-6v-2h1V5.7h-.1L10.835 17H8.137L5.1 5.7H5V15h1v2H1v-2h.25S2 15 2 14.25v-8.5S2 5 1.25 5H1V3h6.634l2.327 8.66h.077L12.386 3H19v2h-.25z"/>`),
		g.Group(children),
	)
}