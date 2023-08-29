package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodePySixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.75 4H3.5a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 1 0v-3h1.75A2.252 2.252 0 0 0 8 6.25A2.252 2.252 0 0 0 5.75 4Zm0 3.5H4V5h1.75A1.252 1.252 0 0 1 7 6.25A1.252 1.252 0 0 1 5.75 7.5Zm8.657-2.71L12 8.16v3.34a.5.5 0 0 1-1 0V8.16L8.593 4.79a.5.5 0 0 1 .294-.78a.5.5 0 0 1 .522.197L11.5 7.14l2.093-2.931a.5.5 0 1 1 .814.582Z"/>`),
		g.Group(children),
	)
}