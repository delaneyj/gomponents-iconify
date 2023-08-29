package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluentTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.724 2.053a.5.5 0 0 0-.448 0l-5 2.5A.5.5 0 0 0 5 5v9.5a.5.5 0 0 0 .243.429l5 3A.5.5 0 0 0 11 17.5v-4.691l4.724-2.362a.5.5 0 0 0 0-.894L11.618 7.5l4.106-2.053a.5.5 0 0 0 0-.894l-5-2.5Z"/>`),
		g.Group(children),
	)
}