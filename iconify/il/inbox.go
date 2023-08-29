package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M736 269q5 9 5 21v244q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V290q0-9 5-21L126 26q13-25 42-25h406q28 0 41 25zm-95 33q3 0 5-3t-1-4l-97-195q-3-6-10-6H203q-7 0-10 6L96 295q-2 2 0 4t4 3h117q14 0 21 13l34 67q6 13 20 13h157q14 0 21-13l33-67q6-13 21-13h117z"/>`),
		g.Group(children),
	)
}