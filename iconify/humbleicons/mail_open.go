package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M21 19V8.588a1 1 0 0 0-.514-.874l-7.03-3.905a3 3 0 0 0-2.913 0L3.514 7.714A1 1 0 0 0 3 8.588V19a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1z"/><path stroke-linecap="round" d="m3 10l7.89 5.26a2 2 0 0 0 2.22 0L21 10M9 14l-6 4m12-4l6 4"/></g>`),
		g.Group(children),
	)
}