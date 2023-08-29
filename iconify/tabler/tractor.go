package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tractor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 15a4 4 0 1 0 8 0a4 4 0 1 0-8 0m4 0v.01M17 17a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-6.5 0H17"/><path d="M20 15.2V11a1 1 0 0 0-1-1h-6l-2-5H5v6.5"/><path d="M18 5h-1a1 1 0 0 0-1 1v4"/></g>`),
		g.Group(children),
	)
}