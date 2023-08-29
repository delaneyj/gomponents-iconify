package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 3h-1a4.002 4.002 0 0 0-3.874 3H6a4 4 0 1 0 0 8h8a4 4 0 0 0 .899-7.899A4.002 4.002 0 0 0 11 3ZM7.676 8l.387-1.501A2.002 2.002 0 0 1 10 5h1c.937 0 1.743.65 1.95 1.549l.28 1.221l1.221.28A2.002 2.002 0 0 1 14 12H6a2 2 0 1 1 0-4h1.676Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}