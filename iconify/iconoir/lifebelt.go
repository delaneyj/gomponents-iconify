package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lifebelt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10Z"/><path d="M8 12a4 4 0 1 0 8 0a4 4 0 0 0-8 0Zm1.235 2.89L5 19m9.765-4.11L19 19m-4.235-9.89L19 5M9.235 9.11L5 5"/></g>`),
		g.Group(children),
	)
}