package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropletHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.578 2.184a1.004 1.004 0 0 0-1.156 0C11.119 2.398 4 7.513 4 13.75C4 18.53 7.364 22 12 22s8-3.468 8-8.246c0-6.241-7.119-11.356-7.422-11.57zM6 13.75c0-4.283 4.395-8.201 6-9.49V20c-3.533 0-6-2.57-6-6.25z"/>`),
		g.Group(children),
	)
}