package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiMacOsWindow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M7 19v-8a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m10 12.01l.01-.011m1.99.011l.01-.011m1.99.011l.01-.011"/><path d="M6.5 16H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v3"/><path stroke-linecap="round" stroke-linejoin="round" d="m5 7.01l.01-.011M7 7.01l.01-.011M9 7.01l.01-.011"/></g>`),
		g.Group(children),
	)
}