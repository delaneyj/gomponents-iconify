package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18.281 14.26L14.08 5.388a2.3 2.3 0 0 0-4.158 0l-.165.349M8.468 8.456L3 20h17"/><path d="m7.5 11l2 2.5l2-2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}