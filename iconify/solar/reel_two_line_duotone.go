package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReelTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="10" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 22h10" opacity=".5"/><path stroke="currentColor" stroke-width="1.5" d="M12 9a3 3 0 1 1 0 6a3 3 0 0 1 0-6Z" opacity=".5"/><path fill="currentColor" d="M19.5 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-13 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM12 4.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm0 13a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z"/></g>`),
		g.Group(children),
	)
}