package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaVideoFill0"><g id="evaVideoFill1"><path id="evaVideoFill2" fill="currentColor" d="M21 7.15a1.7 1.7 0 0 0-1.85.3l-2.15 2V8a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3v-1.45l2.16 2a1.74 1.74 0 0 0 1.16.45a1.68 1.68 0 0 0 .69-.15a1.6 1.6 0 0 0 1-1.48V8.63A1.6 1.6 0 0 0 21 7.15Z"/></g></g>`),
		g.Group(children),
	)
}