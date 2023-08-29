package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Salad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 11h16a1 1 0 0 1 1 1v.5c0 1.5-2.517 5.573-4 6.5v1a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-1c-1.687-1.054-4-5-4-6.5V12a1 1 0 0 1 1-1zm14.5 0c.351-1.017.426-2.236.5-3.714V6h-2.256c-2.83 0-4.616.804-5.64 2.076"/><path d="M5.255 11.008A12.204 12.204 0 0 1 5 9V8h1.755c.98 0 1.801.124 2.479.35M8 8l1-4l4 2.5"/><path d="M13 11v-.5a2.5 2.5 0 1 0-5 0v.5"/></g>`),
		g.Group(children),
	)
}