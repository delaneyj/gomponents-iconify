package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoClipTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.75 4A3.75 3.75 0 0 0 2 7.75v12.5A3.75 3.75 0 0 0 5.75 24h16.5A3.75 3.75 0 0 0 26 20.25V7.75A3.75 3.75 0 0 0 22.25 4H5.75ZM10 10.251a1 1 0 0 1 1.472-.881l7 3.75a1 1 0 0 1 0 1.763l-7 3.75A1 1 0 0 1 10 17.75v-7.5Z"/>`),
		g.Group(children),
	)
}