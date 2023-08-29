package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16a3 3 0 1 1 0 6H7a5 5 0 1 1 4.9-6Zm-2.9-7A6 6 0 0 1 16 4a4.24 4.24 0 0 0 6 6a6 6 0 0 1-3 5.197"/>`),
		g.Group(children),
	)
}