package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21A9 9 0 0 1 5.64 5.64a.74.74 0 0 1 1.06 0a.75.75 0 0 1 0 1.06a7.5 7.5 0 1 0 10.6 10.6a7.48 7.48 0 0 0 0-10.6a.75.75 0 0 1 0-1.06a.74.74 0 0 1 1.06 0A9 9 0 0 1 12 21Z"/><path fill="currentColor" d="M12 12.75a.76.76 0 0 1-.75-.75V4a.75.75 0 0 1 1.5 0v8a.76.76 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}