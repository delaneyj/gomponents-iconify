package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VueJs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 20.5L22.5 4h-4L12 14L5.5 4h-4L12 20.5Z"/><path d="M18.5 4h-4L12 7.5L9.5 4h-4"/></g>`),
		g.Group(children),
	)
}