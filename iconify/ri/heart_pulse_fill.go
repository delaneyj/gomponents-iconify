package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartPulseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 3C19.538 3 22 5.5 22 9c0 7-7.5 11-10 12.5c-1.978-1.186-7.084-3.937-9.131-8.5h4.697l.934-1.556l3 5L13.566 13H17v-2h-4.566l-.934 1.556l-3-5L6.434 11H2.21A9.552 9.552 0 0 1 2 9c0-3.5 2.5-6 5.5-6C9.36 3 11 4 12 5c1-1 2.64-2 4.5-2Z"/>`),
		g.Group(children),
	)
}