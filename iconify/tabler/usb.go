package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 19a2 2 0 1 0 4 0a2 2 0 1 0-4 0m2-2V5.5M7 10v3l5 3m0-1.5l5-2V10m-1 0h2V8h-2z"/><path d="M6 9a1 1 0 1 0 2 0a1 1 0 1 0-2 0m4-3.5h4L12 3z"/></g>`),
		g.Group(children),
	)
}