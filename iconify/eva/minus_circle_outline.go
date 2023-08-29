package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinusCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMinusCircleOutline0"><g id="evaMinusCircleOutline1"><g id="evaMinusCircleOutline2" fill="currentColor"><path d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8Z"/><path d="M15 11H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Z"/></g></g></g>`),
		g.Group(children),
	)
}