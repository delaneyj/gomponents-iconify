package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telecoil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M19 16.5h5m-2.5.5v5m2-21.5L18 6m-6 6L.5 23.5M5.5 9a6.5 6.5 0 0 1 13 0c0 .662-.111 1.32-.328 1.944l-2.85 8.195A3.516 3.516 0 0 1 12 21.5c-.98 0-1.865-.352-2.5-1m6-11.5a3.5 3.5 0 1 0-7 0"/>`),
		g.Group(children),
	)
}