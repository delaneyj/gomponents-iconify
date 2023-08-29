package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsightsTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M11.75 3.5a.75.75 0 0 1 .75.75v15.5a.75.75 0 0 1-1.5 0V4.25a.75.75 0 0 1 .75-.75zm6.5 3.625a.75.75 0 0 1 .75.75V19.75a.75.75 0 0 1-1.5 0V7.875a.75.75 0 0 1 .75-.75zM5.25 11a.75.75 0 0 1 .75.75v8a.75.75 0 0 1-1.5 0v-8a.75.75 0 0 1 .75-.75z" fill="currentColor"/>`),
		g.Group(children),
	)
}