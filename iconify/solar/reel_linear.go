package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReelLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" d="M12 22h10"/><path d="M13.5 7.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm0 9a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm-6-6a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm9 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z"/></g>`),
		g.Group(children),
	)
}