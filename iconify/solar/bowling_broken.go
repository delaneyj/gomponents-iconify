package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlingBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="7" r="1.5" transform="rotate(-90 12 7)"/><circle cx="12" cy="12" r="1.5" transform="rotate(-90 12 12)"/><path d="M8 8a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z"/><path stroke-linecap="round" d="M7 3.338A9.954 9.954 0 0 1 12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12c0-1.821.487-3.53 1.338-5"/></g>`),
		g.Group(children),
	)
}