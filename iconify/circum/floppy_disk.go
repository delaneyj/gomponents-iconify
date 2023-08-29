package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.015 7.015l-4.15-3.39a2.54 2.54 0 0 0-1.58-.56h-9.72a1.5 1.5 0 0 0-1.5 1.5v14.87a1.5 1.5 0 0 0 1.5 1.5h14.87a1.5 1.5 0 0 0 1.5-1.5V8.955a2.507 2.507 0 0 0-.92-1.94Zm-13.45-2.95h5.75v1.37a.5.5 0 0 1-.5.5h-4.75a.5.5 0 0 1-.5-.5Zm0 15.87v-5.93a1.5 1.5 0 0 1 1.5-1.5h7.87a1.5 1.5 0 0 1 1.5 1.5v5.93Zm13.37-.5a.5.5 0 0 1-.5.5h-1v-5.93a2.507 2.507 0 0 0-2.5-2.5h-7.87a2.5 2.5 0 0 0-2.5 2.5v5.93h-1a.5.5 0 0 1-.5-.5V4.565a.5.5 0 0 1 .5-.5h1v1.37a1.5 1.5 0 0 0 1.5 1.5h4.75a1.5 1.5 0 0 0 1.5-1.5v-1.37h.97a1.514 1.514 0 0 1 .95.34l4.14 3.38a1.483 1.483 0 0 1 .56 1.17Z"/>`),
		g.Group(children),
	)
}