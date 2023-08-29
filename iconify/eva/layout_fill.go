package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayoutFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaLayoutFill0"><g id="evaLayoutFill1"><path id="evaLayoutFill2" fill="currentColor" d="M21 8V6a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3v2ZM3 10v8a3 3 0 0 0 3 3h5V10Zm10 0v11h5a3 3 0 0 0 3-3v-8Z"/></g></g>`),
		g.Group(children),
	)
}