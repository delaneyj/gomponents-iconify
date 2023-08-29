package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 11.5V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v1.4m0-.4V8a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v2m0-.1V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v5m0 0v0a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v0"/><path d="M18 11v0a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-4a8 8 0 0 1-8-8a2 2 0 1 1 4 0"/></g>`),
		g.Group(children),
	)
}