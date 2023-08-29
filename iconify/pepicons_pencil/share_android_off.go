package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAndroidOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5 12.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm9-1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm0 14a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z" clip-rule="evenodd"/><path d="m6.754 9.18l-.508-.86l5.5-3.25l.508.86l-5.5 3.25ZM12 14.878l.479-.878l-5.5-3l-.479.878l5.5 3Z"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}