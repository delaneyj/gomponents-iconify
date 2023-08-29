package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21v-9a5 5 0 0 0-5-5H3m9 14v-9a5 5 0 0 1 5-5h4"/><path d="M7 3L3 7l4 4m10-8l4 4l-4 4"/></g>`),
		g.Group(children),
	)
}