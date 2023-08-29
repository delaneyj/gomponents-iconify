package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonClockTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 17.5a6.47 6.47 0 0 1 1.022-3.5h-7.77a2.249 2.249 0 0 0-2.248 2.25v.577c0 .892.318 1.756.898 2.435c1.566 1.834 3.952 2.74 7.098 2.74c.931 0 1.796-.08 2.593-.24A6.475 6.475 0 0 1 11 17.5ZM10 2.005a5 5 0 1 1 0 10a5 5 0 0 1 0-10ZM17.5 12a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11Zm2 5.5h-2V15a.5.5 0 1 0-1 0v3a.5.5 0 0 0 .5.5h2.5a.5.5 0 0 0 0-1Z"/>`),
		g.Group(children),
	)
}