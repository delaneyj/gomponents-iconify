package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleKeyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10a10 10 0 0 1-20 0C2 6.477 6.477 2 12 2zm2 5a3 3 0 0 0-2.98 2.65l-.015.174L11 10l.005.176c.019.319.087.624.197.908l.09.209l-3.5 3.5l-.082.094a1 1 0 0 0 0 1.226l.083.094l1.5 1.5l.094.083a1 1 0 0 0 1.226 0l.094-.083l.083-.094a1 1 0 0 0 0-1.226l-.083-.094l-.792-.793l.585-.585l.793.792l.094.083a1 1 0 0 0 1.403-1.403l-.083-.094l-.792-.793l.792-.792A3 3 0 1 0 14 7zm0 2a1 1 0 1 1 0 2a1 1 0 0 1 0-2z"/></g>`),
		g.Group(children),
	)
}