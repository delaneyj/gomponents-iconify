package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalkieTalkieTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 7.5v3h5v-3h-5ZM8.75 1a.75.75 0 0 1 .75.75V3h7.25A2.25 2.25 0 0 1 19 5.25V14a.75.75 0 0 1-.199.509l-1.301 1.41v3.831A2.25 2.25 0 0 1 15.25 22H8.751a2.25 2.25 0 0 1-2.25-2.25v-3.832L5.2 14.508A.75.75 0 0 1 5 14V5.25A2.25 2.25 0 0 1 7.25 3H8V1.75A.75.75 0 0 1 8.75 1ZM8 6.75v4.5c0 .414.336.75.75.75h6.5a.75.75 0 0 0 .75-.75v-4.5a.75.75 0 0 0-.75-.75h-6.5a.75.75 0 0 0-.75.75Zm1.5 8c0 .414.336.75.75.75h3.5a.75.75 0 0 0 0-1.5h-3.5a.75.75 0 0 0-.75.75Z"/>`),
		g.Group(children),
	)
}