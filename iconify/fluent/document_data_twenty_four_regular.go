package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDataTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.75 11a.75.75 0 0 1 .75.75v6.5a.75.75 0 0 1-1.5 0v-6.5a.75.75 0 0 1 .75-.75Zm4 4.75a.75.75 0 0 0-1.5 0v2.5a.75.75 0 0 0 1.5 0v-2.5Zm2.5-2.75a.75.75 0 0 1 .75.75v4.5a.75.75 0 0 1-1.5 0v-4.5a.75.75 0 0 1 .75-.75ZM12.483 2c.331 0 .65.132.884.366L16.22 5.22a.855.855 0 0 1 .025.025l3.169 3.169A2 2 0 0 1 20 9.828V20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h6.483ZM6 3.5a.5.5 0 0 0-.5.5v16a.5.5 0 0 0 .5.5h12a.5.5 0 0 0 .5-.5V10H14a2 2 0 0 1-2-2V3.5H6Zm7.5 1.121V8a.5.5 0 0 0 .5.5h3.378L13.5 4.621Z"/>`),
		g.Group(children),
	)
}