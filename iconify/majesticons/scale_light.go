package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 6l2 5m0 0h3a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h3z"/><path fill="currentColor" fill-rule="evenodd" d="M2 7a3 3 0 0 1 3-3h2a1 1 0 0 1 1 1v4a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V5a1 1 0 0 1 1-1h2a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V7zm8 9a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-4z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}