package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaskListSquareRtlSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.5A2.5 2.5 0 0 1 4.5 2h7A2.5 2.5 0 0 1 14 4.5v7a2.5 2.5 0 0 1-2.5 2.5h-7A2.5 2.5 0 0 1 2 11.5v-7Zm2.25 6a.5.5 0 0 0 .5.5H7a.5.5 0 0 0 0-1H4.75a.5.5 0 0 0-.5.5Zm.5-4.5a.5.5 0 0 0 0 1H7a.5.5 0 0 0 0-1H4.75Zm6.852 3.146a.5.5 0 0 0-.707 0l-1.147 1.147l-.394-.395a.5.5 0 0 0-.708.707l.748.749a.5.5 0 0 0 .708 0l1.5-1.5a.5.5 0 0 0 0-.708Zm0-3.292a.5.5 0 1 0-.707-.708L9.748 6.293l-.394-.395a.5.5 0 1 0-.708.708l.748.748a.5.5 0 0 0 .708 0l1.5-1.5Z"/>`),
		g.Group(children),
	)
}