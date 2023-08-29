package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 17v-5a3 3 0 0 0-3-3h-5M9 9H6a3 3 0 0 0-3 3v8h17"/><path d="M3 14.803A2.4 2.4 0 0 0 4 15a2.4 2.4 0 0 0 2-1a2.4 2.4 0 0 1 2-1a2.4 2.4 0 0 1 2 1a2.4 2.4 0 0 0 2 1a2.4 2.4 0 0 0 2-1m4 0a2.4 2.4 0 0 0 2 1c.35.007.692-.062 1-.197M10.172 6.188c.07-.158.163-.31.278-.451L12 4l1.465 1.638a2 2 0 0 1-.65 3.19M3 3l18 18"/></g>`),
		g.Group(children),
	)
}