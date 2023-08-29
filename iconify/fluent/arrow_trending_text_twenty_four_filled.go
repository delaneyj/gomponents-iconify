package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTrendingTextTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21 4a1 1 0 0 0-1-1h-5a1 1 0 1 0 0 2h2.586L12.5 10.086l-1.793-1.793a1 1 0 0 0-1.414 0l-6 6a1 1 0 1 0 1.414 1.414L10 10.414l1.793 1.793a1 1 0 0 0 .564.283c.296-.194.628-.338.985-.418L19 6.414V9a1 1 0 1 0 2 0V4Zm-7 9a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2h-7Zm0 3.5a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5Zm0 3a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}