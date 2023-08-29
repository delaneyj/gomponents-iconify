package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeShop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M18.608 8.5h.892V9c-.577 2.831-1 5.714-1 8.604v.896h-.911m1.019-10c.24-1.508.52-3.01.84-4.506c.034-.162.052-.328.052-.494h-15c0 .166.018.332.052.494c.32 1.496.6 2.998.84 4.506m13.216 0H5.392m12.197 10a88.458 88.458 0 0 0-.089 3.964V23.5h-11v-1.036c0-1.323-.03-2.644-.089-3.964m11.178 0H6.41m-1.019-10H4.5V9c.577 2.831 1 5.714 1 8.604v.896h.911M2.5 3.5V3l.586-.293A3.817 3.817 0 0 0 5 .5h14a3.817 3.817 0 0 0 1.914 2.207L21.5 3v.5h-19Z"/>`),
		g.Group(children),
	)
}