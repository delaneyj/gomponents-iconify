package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataScatterTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.5a.5.5 0 0 0-1 0v15a.5.5 0 0 0 .5.5h15a.5.5 0 0 0 0-1H3V2.5ZM7.5 10a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM17 5.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0ZM12.5 15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/>`),
		g.Group(children),
	)
}