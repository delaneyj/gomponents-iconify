package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Humbleicon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M8 4a2 2 0 0 1 2 2v11a2 2 0 1 1-4 0V6a2 2 0 0 1 2-2zm8 8a2 2 0 0 1 2 2v3a2 2 0 1 1-4 0v-3a2 2 0 0 1 2-2z"/><circle cx="16" cy="6" r="2"/></g>`),
		g.Group(children),
	)
}