package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableMoveRightTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24.25 3a.75.75 0 0 1 .75.75v20.5a.75.75 0 0 1-1.5 0V3.75a.75.75 0 0 1 .75-.75Zm-6.163 7.442a1.744 1.744 0 0 0-.419.558H11v6h6.668a1.745 1.745 0 0 0 .832.832V24a1 1 0 0 1-1 1H6.75A3.75 3.75 0 0 1 3 21.25V6.75A3.75 3.75 0 0 1 6.75 3H17.5a1 1 0 0 1 1 1v6.168c-.146.07-.285.161-.413.274ZM17 9.5v-5h-6v5h6Zm-7.5 0v-5H6.75A2.25 2.25 0 0 0 4.5 6.75V9.5h5Zm0 1.5h-5v6h5v-6Zm-5 10.25a2.25 2.25 0 0 0 2.25 2.25H9.5v-5h-5v2.75ZM11 23.5h6v-5h-6v5Zm7.69-7.748l.89-1.002H16a.75.75 0 0 1 0-1.5h3.58l-.89-1.002a.75.75 0 0 1 1.12-.996l2 2.25a.75.75 0 0 1 0 .996l-2 2.25a.75.75 0 0 1-1.12-.996Z"/>`),
		g.Group(children),
	)
}