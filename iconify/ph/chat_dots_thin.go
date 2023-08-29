package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatDotsThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 52H40a12 12 0 0 0-12 12v160a11.89 11.89 0 0 0 6.93 10.88A12.17 12.17 0 0 0 40 236a11.89 11.89 0 0 0 7.69-2.83l.06-.06l32.14-28.17a4 4 0 0 1 2.61-.94H216a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12Zm4 140a4 4 0 0 1-4 4H82.5a12.1 12.1 0 0 0-7.79 2.87l-32.16 28.2A4 4 0 0 1 36 224V64a4 4 0 0 1 4-4h176a4 4 0 0 1 4 4Zm-84-64a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm-44 0a8 8 0 1 1-8-8a8 8 0 0 1 8 8Zm88 0a8 8 0 1 1-8-8a8 8 0 0 1 8 8Z"/>`),
		g.Group(children),
	)
}