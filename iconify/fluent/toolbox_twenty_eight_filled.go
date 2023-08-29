package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolboxTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.002 6.247v1.748H4.25A2.25 2.25 0 0 0 2 10.245V14h5.506v-1.25a.75.75 0 0 1 1.5 0V14h10v-1.25a.75.75 0 0 1 1.5 0V14h5.497v-3.755a2.25 2.25 0 0 0-2.25-2.25h-3.75V6.247a2.25 2.25 0 0 0-2.25-2.25h-7.5a2.25 2.25 0 0 0-2.25 2.25Zm2.25-.75h7.501a.75.75 0 0 1 .75.75v1.748h-9V6.247a.75.75 0 0 1 .75-.75ZM26.003 15.5h-5.497v1.25a.75.75 0 0 1-1.5 0V15.5h-10v1.25a.75.75 0 0 1-1.5 0V15.5H2v6.245a2.25 2.25 0 0 0 2.25 2.25h19.503a2.25 2.25 0 0 0 2.25-2.25V15.5Z"/>`),
		g.Group(children),
	)
}