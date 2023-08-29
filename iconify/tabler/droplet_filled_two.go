package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropletFilledTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.8 11a6 6 0 1 0 10.396 0l-5.197-8l-5.2 8zM6 14h12M7.305 17.695L11 14"/><path d="M10.26 19.74L16 14l-5.74 5.74z"/></g>`),
		g.Group(children),
	)
}