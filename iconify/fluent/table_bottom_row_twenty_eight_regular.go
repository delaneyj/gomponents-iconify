package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableBottomRowTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.75 3A3.75 3.75 0 0 0 3 6.75v14.5A3.75 3.75 0 0 0 6.75 25h14.5A3.75 3.75 0 0 0 25 21.25V6.75A3.75 3.75 0 0 0 21.25 3H6.75ZM4.5 21.25V18.5h5v5H6.75a2.25 2.25 0 0 1-2.25-2.25ZM11 23.5v-5h6v5h-6ZM4.5 17V6.75A2.25 2.25 0 0 1 6.75 4.5h14.5a2.25 2.25 0 0 1 2.25 2.25V17h-19Zm14 6.5v-5h5v2.75a2.25 2.25 0 0 1-2.25 2.25H18.5Z"/>`),
		g.Group(children),
	)
}