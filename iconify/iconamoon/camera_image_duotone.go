package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraImageDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M21 6H3v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6Zm-9 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5Z"/><circle cx="12" cy="12" r="3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 2h2"/></g>`),
		g.Group(children),
	)
}