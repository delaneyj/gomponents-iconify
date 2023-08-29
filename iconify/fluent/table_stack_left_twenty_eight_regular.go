package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableStackLeftTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 24.25a.75.75 0 0 0 1.5 0V3.75a.75.75 0 0 0-1.5 0v20.5Zm22-17.5A3.75 3.75 0 0 0 21.25 3H10.5a1 1 0 0 0-1 1v20a1 1 0 0 0 1 1h10.75A3.75 3.75 0 0 0 25 21.25V6.75ZM11 18.5h6v5h-6v-5Zm6-1.5h-6v-6h6v6Zm1.5 1.5h5v2.75a2.25 2.25 0 0 1-2.25 2.25H18.5v-5Zm5-1.5h-5v-6h5v6Zm0-10.25V9.5h-5v-5h2.75a2.25 2.25 0 0 1 2.25 2.25ZM17 4.5v5h-6v-5h6Z"/>`),
		g.Group(children),
	)
}