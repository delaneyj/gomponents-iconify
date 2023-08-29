package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraImageThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5Z"/><circle cx="12" cy="12" r="3" stroke-linecap="round"/><path stroke-linecap="round" d="M17 2h2"/></g>`),
		g.Group(children),
	)
}