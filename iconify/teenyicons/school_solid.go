package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchoolSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 8a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1Z"/><path fill="currentColor" fill-rule="evenodd" d="m12 6.191l-4-2V3h3V0H7v4.191l-4 2V8H1v6H0v1h6v-4h3v4h6v-1h-1V8h-2V6.191ZM13 14V9h-1v5h1ZM3 14H2V9h1v5Zm3-5.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0Z" clip-rule="evenodd"/><path fill="currentColor" d="M8 15v-3H7v3h1Z"/>`),
		g.Group(children),
	)
}