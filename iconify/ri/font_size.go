package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontSize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.245 15h-6.49l-2 5H.6L7 4h2l6.4 16h-2.155l-2-5Zm-.8-2L8 6.885L5.554 13h4.891ZM21 12.535V12h2v8h-2v-.535a4 4 0 1 1 0-6.93ZM19 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}