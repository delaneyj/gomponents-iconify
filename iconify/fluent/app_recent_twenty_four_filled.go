package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppRecentTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 19.75a2.25 2.25 0 0 1-2.096 2.245L15.75 22h-7.5a2.25 2.25 0 0 1-2.245-2.096L6 19.75V4.25a2.25 2.25 0 0 1 2.096-2.245L8.25 2h7.5a2.25 2.25 0 0 1 2.245 2.096L18 4.25v15.5ZM19 5h.75a2.25 2.25 0 0 1 2.245 2.096L22 7.25v9.5a2.25 2.25 0 0 1-2.096 2.245L19.75 19H19V5ZM5 19h-.75a2.25 2.25 0 0 1-2.245-2.096L2 16.75v-9.5a2.25 2.25 0 0 1 2.096-2.245L4.25 5H5v14Z"/>`),
		g.Group(children),
	)
}