package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DonutBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="3"/><path d="M2 13s2.2 2 4 2c1.212 0 2.606-.908 3.387-1.5m4.613.724c.471.415 1.088.776 1.805.776c1.69 0 1.69-2 3.38-2c1.077 0 1.925.814 2.399 1.403"/><path stroke-linecap="round" d="M14.5 7L16 5m3 2l1-1m-8-1l-1-1m-.5 3l-1.366.366m7.516 1.611l.066 1.412m3.962-.304L19 11.563M7 5L6 4m.792 5.144l-.585-1.288m-.542 4.786L6.5 11.5m-2.817-1.15l-.079-1.412M7 20.662A9.955 9.955 0 0 0 12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12c0 1.821.487 3.53 1.338 5"/></g>`),
		g.Group(children),
	)
}