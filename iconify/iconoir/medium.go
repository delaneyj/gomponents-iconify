package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm8 0c1.105 0 2-1.79 2-4s-.895-4-2-4s-2 1.79-2 4s.895 4 2 4Zm5 0c.552 0 1-1.79 1-4s-.448-4-1-4s-1 1.79-1 4s.448 4 1 4Z"/>`),
		g.Group(children),
	)
}