package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCopy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M10 4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2h-4zm0-2a3.001 3.001 0 0 0-2.83 2H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-4h-8.586l1.293 1.293a1 1 0 0 1-1.414 1.414l-3-3a1 1 0 0 1 0-1.414l3-3a1 1 0 1 1 1.414 1.414L12.414 13H21V7a3 3 0 0 0-3-3h-1.17A3.001 3.001 0 0 0 14 2h-4z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}