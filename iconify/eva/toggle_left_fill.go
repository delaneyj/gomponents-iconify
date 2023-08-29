package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaToggleLeftFill0"><g id="evaToggleLeftFill1"><g id="evaToggleLeftFill2" fill="currentColor"><path d="M15 5H9a7 7 0 0 0 0 14h6a7 7 0 0 0 0-14ZM9 15a3 3 0 1 1 3-3a3 3 0 0 1-3 3Z"/><path d="M9 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/></g></g></g>`),
		g.Group(children),
	)
}