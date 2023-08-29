package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 1.5A1.5 1.5 0 0 1 1.5 0h12A1.5 1.5 0 0 1 15 1.5V9H9.5a.5.5 0 0 0-.5.5V15H1.5A1.5 1.5 0 0 1 0 13.5v-12ZM12 4H3V3h9v1Z" clip-rule="evenodd"/><path fill="currentColor" d="M10 15h.086a1.5 1.5 0 0 0 1.06-.44l3.415-3.414a1.5 1.5 0 0 0 .439-1.06V10h-5v5Z"/>`),
		g.Group(children),
	)
}