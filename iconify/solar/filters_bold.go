package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiltersBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M18 8A6 6 0 1 1 6 8a6 6 0 0 1 12 0Z"/><path d="M5.033 10.783a6 6 0 1 0 8.92 4.46a7.503 7.503 0 0 1-8.92-4.46Zm10.354 3.911c.074.424.113.86.113 1.306c0 2.09-.855 3.982-2.235 5.342a6 6 0 0 0 5.702-10.558a7.527 7.527 0 0 1-3.58 3.91Z"/></g>`),
		g.Group(children),
	)
}