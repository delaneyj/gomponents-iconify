package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MountainF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.797 10.402c-2.038-.98-5.317-1.604-8.3-1.604c-1.828 0-2.951-3.163-3.264-5.235l1.182-2.007a3 3 0 0 1 5.17 0l5.212 8.846zm1.673 2.839l.14.236A3 3 0 0 1 17.024 18H2.975A3 3 0 0 1 .39 13.477l4.517-7.664c.417 2.76 2.144 4.985 4.59 4.985c2.747 0 6.71.536 9.973 2.443z"/>`),
		g.Group(children),
	)
}