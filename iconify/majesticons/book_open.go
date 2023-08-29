package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 4.85c-2.195-1.022-5.52-1.565-8.633.979A1 1 0 0 0 2 6.603V19a1 1 0 0 0 1.633.774c2.736-2.236 5.734-1.31 7.367-.255V4.849zm2 0v14.669c1.633-1.056 4.63-1.981 7.367.255A1 1 0 0 0 22 19V6.603a1 1 0 0 0-.367-.774C18.52 3.285 15.195 3.828 13 4.849z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}