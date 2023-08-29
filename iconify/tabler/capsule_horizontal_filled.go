package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CapsuleHorizontalFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M15 5H9a7 7 0 1 0 0 14h6a7 7 0 0 0 7-7l-.007-.303A7 7 0 0 0 15 5z"/></g>`),
		g.Group(children),
	)
}