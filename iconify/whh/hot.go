package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M416 1024q-113 0-209-55.5T55.5 817T0 608q0-44 13-79.5T45 465t38-61t32-87t13-125q0 10 10.5 19.5t25 19t32 33.5t28.5 56q16 49 52 88.5t76 39.5q52 0 88.5-59t54-159T512 0q15 44 46 92t65.5 87t72.5 89t67.5 97.5t49 111T832 608q0 113-55.5 209T625 968.5T416 1024zm128-640q-51 136-96 192q-19 24-47.5 43t-52 30.5t-45.5 26t-34.5 37T256 768q0 61 47.5 94.5T416 896q96 0 160-50.5T640 708q0-67-11-115.5t-25-72t-32-60t-28-76.5z"/>`),
		g.Group(children),
	)
}