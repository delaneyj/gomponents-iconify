package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func S(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 14a4 4 0 0 0 0 8h4a8 8 0 1 1 0 16h-4a8.003 8.003 0 0 1-7.544-5.334a2 2 0 0 1 3.771-1.332A4.002 4.002 0 0 0 22 34h4a4 4 0 0 0 0-8h-4a8 8 0 1 1 0-16h4a8.003 8.003 0 0 1 7.544 5.334a2 2 0 0 1-3.771 1.332A4.002 4.002 0 0 0 26 14h-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}