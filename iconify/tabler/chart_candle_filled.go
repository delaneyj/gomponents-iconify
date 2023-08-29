package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartCandleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M6 3a1 1 0 0 1 .993.883L7 4v1a2 2 0 0 1 1.995 1.85L9 7v3a2 2 0 0 1-1.85 1.995L7 12v8a1 1 0 0 1-1.993.117L5 20v-8a2 2 0 0 1-1.995-1.85L3 10V7a2 2 0 0 1 1.85-1.995L5 5V4a1 1 0 0 1 1-1zm6 0a1 1 0 0 1 .993.883L13 4v9a2 2 0 0 1 1.995 1.85L15 15v3a2 2 0 0 1-1.85 1.995L13 20a1 1 0 0 1-1.993.117L11 20l-.15-.005a2 2 0 0 1-1.844-1.838L9 18v-3a2 2 0 0 1 1.85-1.995L11 13V4a1 1 0 0 1 1-1zm6 0a1 1 0 0 1 .993.883L19 4a2 2 0 0 1 1.995 1.85L21 6v4a2 2 0 0 1-1.85 1.995L19 12v8a1 1 0 0 1-1.993.117L17 20v-8a2 2 0 0 1-1.995-1.85L15 10V6a2 2 0 0 1 1.85-1.995L17 4a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}