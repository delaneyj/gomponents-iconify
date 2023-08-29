package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 3A2.25 2.25 0 0 0 4 5.25v13.5A2.25 2.25 0 0 0 6.25 21h2.5A2.25 2.25 0 0 0 11 18.75V5.25A2.25 2.25 0 0 0 8.75 3h-2.5ZM5.5 5.25a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 .75.75v13.5a.75.75 0 0 1-.75.75h-2.5a.75.75 0 0 1-.75-.75V5.25ZM15.25 3A2.25 2.25 0 0 0 13 5.25v13.5A2.25 2.25 0 0 0 15.25 21h2.5A2.25 2.25 0 0 0 20 18.75V5.25A2.25 2.25 0 0 0 17.75 3h-2.5Zm-.75 2.25a.75.75 0 0 1 .75-.75h2.5a.75.75 0 0 1 .75.75v13.5a.75.75 0 0 1-.75.75h-2.5a.75.75 0 0 1-.75-.75V5.25Z"/>`),
		g.Group(children),
	)
}