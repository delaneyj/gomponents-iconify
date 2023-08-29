package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M10 2a8 8 0 1 1 0 16a8 8 0 0 1 0-16zm3.5 7.5h-7l-.09.008a.5.5 0 0 0 .09.992h7l.09-.008a.5.5 0 0 0-.09-.992z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}