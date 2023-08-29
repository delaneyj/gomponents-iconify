package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddKeyframes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12h3m3 0H5m0 0V9m0 3v3m1.25-9l.245-.28a2 2 0 0 1 3.01 0l4.343 4.963a2 2 0 0 1 0 2.634L9.505 18.28a2 2 0 0 1-3.01 0L6.25 18"/><path d="m13 19l4.884-5.698a2 2 0 0 0 0-2.604L13 5"/><path d="m17 19l4.884-5.698a2 2 0 0 0 0-2.604L17 5"/></g>`),
		g.Group(children),
	)
}