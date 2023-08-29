package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 17a2 2 0 1 0 4 0a2 2 0 1 0-4 0m10.584-1.412a2 2 0 0 0 2.828 2.83"/><path d="M5 17H3v-6l2-5h1m4 0h4l4 5h1a2 2 0 0 1 2 2v4m-6 0H9m-6-6h8m4 0h3m-6-3V6M3 3l18 18"/></g>`),
		g.Group(children),
	)
}