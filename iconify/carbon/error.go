package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Error(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 16A14 14 0 1 0 16 2A14 14 0 0 0 2 16Zm23.15 7.75L8.25 6.85a12 12 0 0 1 16.9 16.9ZM8.24 25.16a12 12 0 0 1-1.4-16.89l16.89 16.89a12 12 0 0 1-15.49 0Z"/>`),
		g.Group(children),
	)
}