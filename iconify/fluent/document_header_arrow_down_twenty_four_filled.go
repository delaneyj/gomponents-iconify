package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentHeaderArrowDownTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.75 2a2.25 2.25 0 0 1 2.245 2.096L20 4.25v7.248A6.5 6.5 0 0 0 12.81 22H6.25a2.25 2.25 0 0 1-2.245-2.096L4 19.75V4.25a2.25 2.25 0 0 1 2.096-2.245L6.25 2h11.5ZM8.505 5.005a1.5 1.5 0 0 0 0 3h7a1.5 1.5 0 0 0 0-3h-7ZM23 17.5a5.5 5.5 0 1 1-11 0a5.5 5.5 0 0 1 11 0Zm-5-3a.5.5 0 0 0-1 0v4.793l-1.646-1.647a.5.5 0 0 0-.708.708l2.5 2.5a.5.5 0 0 0 .708 0l2.5-2.5a.5.5 0 0 0-.708-.708L18 19.293V14.5Z"/>`),
		g.Group(children),
	)
}