package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiceThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M3 20.4V3.6a.6.6 0 0 1 .6-.6h16.8a.6.6 0 0 1 .6.6v16.8a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6Z"/><path fill="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 8a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm4.5 4.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Zm4.5 4.5a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1Z"/></g>`),
		g.Group(children),
	)
}