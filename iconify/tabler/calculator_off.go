package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.823 19.824A2 2 0 0 1 18 21H6a2 2 0 0 1-2-2V5c0-.295.064-.575.178-.827M7 3h11a2 2 0 0 1 2 2v11"/><path d="M10 10H9a1 1 0 0 1-1-1V8m3-1h4a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1m-6 4v.01m4-.01v.01M8 17v.01m4-.01v.01m4-.01v.01M3 3l18 18"/></g>`),
		g.Group(children),
	)
}