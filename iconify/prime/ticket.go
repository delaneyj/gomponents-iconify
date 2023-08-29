package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 18.75H5A1.76 1.76 0 0 1 3.25 17v-2.5a.76.76 0 0 1 .75-.75a1.75 1.75 0 0 0 0-3.5a.76.76 0 0 1-.75-.75V7A1.76 1.76 0 0 1 5 5.25h14A1.76 1.76 0 0 1 20.75 7v2.5a.76.76 0 0 1-.75.75a1.75 1.75 0 0 0 0 3.5a.76.76 0 0 1 .75.75V17A1.76 1.76 0 0 1 19 18.75ZM4.75 15.16V17a.25.25 0 0 0 .25.25h14a.25.25 0 0 0 .25-.25v-1.84a3.25 3.25 0 0 1 0-6.32V7a.25.25 0 0 0-.25-.25H5a.25.25 0 0 0-.25.25v1.84a3.25 3.25 0 0 1 0 6.32Z"/>`),
		g.Group(children),
	)
}