package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableStackAboveTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.25 4.5a.75.75 0 0 0 0-1.5H3.75a.75.75 0 0 0 0 1.5h16.5Zm0 4a.75.75 0 0 1 .75.75v8.5A3.25 3.25 0 0 1 17.75 21H6.25A3.25 3.25 0 0 1 3 17.75v-8.5a.75.75 0 0 1 .75-.75h16.5ZM14 10h-4v4h4v-4Zm-5.5 0h-4v4h4v-4Zm0 5.5h-4v2.25c0 .966.784 1.75 1.75 1.75H8.5v-4Zm5.5 0h-4v4h4v-4Zm1.5 4h2.25a1.75 1.75 0 0 0 1.75-1.75V15.5h-4v4Zm4-9.5h-4v4h4v-4Z"/>`),
		g.Group(children),
	)
}