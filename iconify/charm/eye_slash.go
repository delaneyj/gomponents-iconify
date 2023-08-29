package charm

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.75 3.75c3.5.5 5.5 4.25 5.5 4.25s-.5 1.25-1.5 2.25m-2.5 1.5c-6 2-8.5-3.75-8.5-3.75s.5-1.75 3-3.25"/><path fill="currentColor" d="M8.625 9.083a1.25 1.25 0 0 1-1.649-.366a1.25 1.25 0 0 1 .22-1.675L8 8z"/><path d="m3.75 1.75l8.5 12.5"/></g>`),
		g.Group(children),
	)
}