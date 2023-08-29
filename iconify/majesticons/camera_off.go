package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l18 18"/><path fill="currentColor" fill-rule="evenodd" d="M5.172 6H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14c.35 0 .688-.06 1-.17l-5.087-5.089a4 4 0 1 1-5.654-5.654L5.172 6zm5.502 5.503a2 2 0 1 0 2.823 2.823l-2.823-2.823zM22 17.172V9a3 3 0 0 0-3-3h-.93a1 1 0 0 1-.832-.445l-.812-1.22A3 3 0 0 0 13.93 3h-3.86a3 3 0 0 0-1.708.534L22 17.172z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}