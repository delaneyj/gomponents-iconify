package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OptionsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="9" cy="16" r="2" fill="currentColor" opacity=".16"/><circle cx="15" cy="8" r="2" fill="currentColor" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h9m4 0h3m-9 8h9M4 16h3"/><circle cx="9" cy="16" r="2" stroke="currentColor" stroke-width="2"/><circle cx="15" cy="8" r="2" stroke="currentColor" stroke-width="2"/></g>`),
		g.Group(children),
	)
}