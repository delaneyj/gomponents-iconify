package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3h4v4m0-4L6 18m-3 0h3v3m10.5-1c1.576-1.576 2.5-4.095 2.5-6.5C19 8.69 15.31 5 10.5 5C8.085 5 5.578 5.913 4 7.5L16.5 20z"/>`),
		g.Group(children),
	)
}