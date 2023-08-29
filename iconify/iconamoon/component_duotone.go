package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l3 3l-3 3l-3-3l3-3Zm0 12l3 3l-3 3l-3-3l3-3Zm6-6l3 3l-3 3l-3-3l3-3ZM6 9l3 3l-3 3l-3-3l3-3Z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m12 3l3 3l-3 3l-3-3l3-3Zm0 12l3 3l-3 3l-3-3l3-3Zm6-6l3 3l-3 3l-3-3l3-3ZM6 9l3 3l-3 3l-3-3l3-3Z"/></g>`),
		g.Group(children),
	)
}