package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19.5 12.572l-.536.53m-7.91 5.96L4.5 12.573A5 5 0 1 1 12 6.006a5 5 0 1 1 7.5 6.572M20 21l2-2l-2-2m-3 0l-2 2l2 2"/>`),
		g.Group(children),
	)
}