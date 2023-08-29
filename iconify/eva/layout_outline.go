package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaLayoutOutline0"><g id="evaLayoutOutline1"><path id="evaLayoutOutline2" fill="currentColor" d="M18 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3ZM6 5h12a1 1 0 0 1 1 1v2H5V6a1 1 0 0 1 1-1ZM5 18v-8h6v9H6a1 1 0 0 1-1-1Zm13 1h-5v-9h6v8a1 1 0 0 1-1 1Z"/></g></g>`),
		g.Group(children),
	)
}