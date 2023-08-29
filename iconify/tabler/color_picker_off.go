package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorPickerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 7l6 6m-5-5l3.699-3.699a1 1 0 0 1 1.4 0l2.6 2.6a1 1 0 0 1 0 1.4l-3.702 3.702m-2 2l-6 6h-4v-4l6-6M3 3l18 18"/>`),
		g.Group(children),
	)
}