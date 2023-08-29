package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReelTwoBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 22h10"/><path stroke="currentColor" stroke-width="1.5" d="M12 9a3 3 0 1 1 0 6a3 3 0 0 1 0-6Z"/><path fill="currentColor" d="M19.5 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-13 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM12 4.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm0 13a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5"/></g>`),
		g.Group(children),
	)
}