package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsAddInSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.25 1.75a.75.75 0 0 0-1.5 0v1h-1a.75.75 0 0 0 0 1.5h1v1a.75.75 0 0 0 1.5 0v-1h1a.75.75 0 0 0 0-1.5h-1v-1ZM4 2a2 2 0 0 0-2 2v7.75a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-2.5a2 2 0 0 0-2-2H8.75V4a2 2 0 0 0-2-2H4Zm3.25 6.75v3.5H4a.5.5 0 0 1-.5-.5v-3h3.75Zm0-4.75v3.25H3.5V4a.5.5 0 0 1 .5-.5h2.75a.5.5 0 0 1 .5.5ZM12 12.25H8.75v-3.5H12a.5.5 0 0 1 .5.5v2.5a.5.5 0 0 1-.5.5Z"/>`),
		g.Group(children),
	)
}