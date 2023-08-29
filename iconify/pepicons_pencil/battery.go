package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Battery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><rect width="2" height="5" x="16" y="7.5" rx=".5"/><path d="M4 7.5h2.5v5H4v-5Zm3.25 0h2.5v5h-2.5v-5Zm3.25 0H13v5h-2.5z"/><path fill-rule="evenodd" d="M14 5.5H3A1.5 1.5 0 0 0 1.5 7v6A1.5 1.5 0 0 0 3 14.5h11a1.5 1.5 0 0 0 1.5-1.5V7A1.5 1.5 0 0 0 14 5.5ZM2.5 7a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-.5.5H3a.5.5 0 0 1-.5-.5V7Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}