package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchoolTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="10" r="1"/><path d="M22 20V8h-4l-6-4l-6 4H2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2ZM6 17v.01M6 13v.01M18 17v.01M18 13v.01"/><path d="M14 22v-5a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v5"/></g>`),
		g.Group(children),
	)
}