package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAppstore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12a9 9 0 1 0 18 0a9 9 0 1 0-18 0m5 4l1.106-1.99m1.4-2.522L13 7m-6 7h5m2.9 0H17m-1 2l-2.51-4.518m-1.487-2.677l-1-1.805"/>`),
		g.Group(children),
	)
}