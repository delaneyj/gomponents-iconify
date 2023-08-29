package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneScreenTimeTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.5 12a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11ZM13.75 2A2.25 2.25 0 0 1 16 4.25v6.924A6.503 6.503 0 0 0 11.02 18H8.75a.75.75 0 0 0-.102 1.493l.102.007h2.5l.062-.003A6.496 6.496 0 0 0 12.81 22H6.25A2.25 2.25 0 0 1 4 19.75V4.25A2.25 2.25 0 0 1 6.25 2h7.5Zm2.75 12a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h3.002a.5.5 0 0 0 0-1H17v-3.5a.5.5 0 0 0-.5-.5Z"/>`),
		g.Group(children),
	)
}