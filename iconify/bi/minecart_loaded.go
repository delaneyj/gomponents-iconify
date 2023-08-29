package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinecartLoaded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M4 15a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm8-1a1 1 0 1 1 0-2a1 1 0 0 1 0 2zm0 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4zM.115 3.18A.5.5 0 0 1 .5 3h15a.5.5 0 0 1 .491.592l-1.5 8A.5.5 0 0 1 14 12H2a.5.5 0 0 1-.491-.408l-1.5-8a.5.5 0 0 1 .106-.411zm.987.82l1.313 7h11.17l1.313-7H1.102z"/><path fill-rule="evenodd" d="M6 1a2.498 2.498 0 0 1 4 0c.818 0 1.545.394 2 1c.67 0 1.552.57 2 1h-2c-.314 0-.611-.15-.8-.4c-.274-.365-.71-.6-1.2-.6c-.314 0-.611-.15-.8-.4a1.497 1.497 0 0 0-2.4 0c-.189.25-.486.4-.8.4c-.507 0-.955.251-1.228.638c-.09.13-.194.25-.308.362H3c.13-.147.401-.432.562-.545a1.63 1.63 0 0 0 .393-.393A2.498 2.498 0 0 1 6 1z"/></g>`),
		g.Group(children),
	)
}