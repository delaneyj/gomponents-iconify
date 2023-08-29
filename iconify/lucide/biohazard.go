package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Biohazard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="11.9" r="2"/><path d="M6.7 3.4c-.9 2.5 0 5.2 2.2 6.7C6.5 9 3.7 9.6 2 11.6m6.9-1.5l1.4.8m7-7.5c.9 2.5 0 5.2-2.2 6.7c2.4-1.2 5.2-.6 6.9 1.5m-6.9-1.5l-1.4.8m3 9.9c-2.6-.4-4.6-2.6-4.7-5.3c-.2 2.6-2.1 4.8-4.7 5.2m4.7-6.8v1.6m1.5-10.1c-1-.2-2-.2-3 0m6.5 11c.7-.7 1.2-1.6 1.5-2.5m-13 0c.3.9.8 1.8 1.5 2.5"/></g>`),
		g.Group(children),
	)
}