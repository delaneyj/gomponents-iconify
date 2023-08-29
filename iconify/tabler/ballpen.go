package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ballpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 6l7 7l-4 4"/><path d="M5.828 18.172a2.828 2.828 0 0 0 4 0L20.414 7.586a2 2 0 0 0 0-2.829l-1.171-1.171a2 2 0 0 0-2.829 0L5.828 14.172a2.828 2.828 0 0 0 0 4zM4 20l1.768-1.768"/></g>`),
		g.Group(children),
	)
}