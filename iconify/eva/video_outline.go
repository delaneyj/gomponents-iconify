package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaVideoOutline0"><g id="evaVideoOutline1"><path id="evaVideoOutline2" fill="currentColor" d="M21 7.15a1.7 1.7 0 0 0-1.85.3l-2.15 2V8a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3v-1.45l2.16 2a1.74 1.74 0 0 0 1.16.45a1.68 1.68 0 0 0 .69-.15a1.6 1.6 0 0 0 1-1.48V8.63A1.6 1.6 0 0 0 21 7.15ZM15 16a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1Zm5-1.4L17.19 12L20 9.4Z"/></g></g>`),
		g.Group(children),
	)
}