package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableFreezeColumnAndRowTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.25 3A3.75 3.75 0 0 1 25 6.75v14.5A3.75 3.75 0 0 1 21.25 25H6.75A3.75 3.75 0 0 1 3 21.25V6.75A3.75 3.75 0 0 1 6.75 3h14.5ZM4.5 6.75V9.5h14v14h2.75a2.25 2.25 0 0 0 2.25-2.25V6.75a2.25 2.25 0 0 0-2.25-2.25H6.75A2.25 2.25 0 0 0 4.5 6.75ZM11 17h6v-6h-6v6Zm6 6.5v-5h-6v5h6ZM9.5 11h-5v6h5v-6Zm-5 7.5v2.75a2.25 2.25 0 0 0 2.25 2.25H9.5v-5h-5Z"/>`),
		g.Group(children),
	)
}