package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CouchTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.75 4h-9.5A2.75 2.75 0 0 0 4.5 6.75V8a3.5 3.5 0 0 1 3.465 3h8.07A3.5 3.5 0 0 1 19.5 8V6.75A2.75 2.75 0 0 0 16.75 4Zm.25 7.5a2.5 2.5 0 0 1 5 0v2.75a2.75 2.75 0 0 1-2.5 2.739v1.261a.75.75 0 0 1-1.5 0V17H6v1.25a.75.75 0 0 1-1.5 0v-1.261A2.75 2.75 0 0 1 2 14.25V11.5a2.5 2.5 0 0 1 5 0a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5Z"/>`),
		g.Group(children),
	)
}