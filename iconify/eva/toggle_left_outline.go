package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaToggleLeftOutline0"><g id="evaToggleLeftOutline1"><g id="evaToggleLeftOutline2" fill="currentColor"><path d="M15 5H9a7 7 0 0 0 0 14h6a7 7 0 0 0 0-14Zm0 12H9A5 5 0 0 1 9 7h6a5 5 0 0 1 0 10Z"/><path d="M9 9a3 3 0 1 0 3 3a3 3 0 0 0-3-3Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g></g>`),
		g.Group(children),
	)
}