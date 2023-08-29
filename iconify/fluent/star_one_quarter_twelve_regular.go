package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOneQuarterTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.718 1.546a.8.8 0 0 0-1.435 0L4.172 3.798l-2.486.361a.798.798 0 0 0-.444 1.365l1.8 1.753l-.425 2.476a.837.837 0 0 0-.005.254a.801.801 0 0 0 1.165.59l2.224-1.17l2.223 1.17a.8.8 0 0 0 1.161-.844L8.96 7.277l1.8-1.753a.8.8 0 0 0-.444-1.365L7.83 3.8L6.718 1.545ZM5 8.824V4.376a.799.799 0 0 0 .022-.041L6 2.352l.979 1.983a.8.8 0 0 0 .602.438l2.19.318l-1.585 1.544a.8.8 0 0 0-.23.708l.374 2.18l-1.958-1.03a.8.8 0 0 0-.744 0L5 8.825Z"/>`),
		g.Group(children),
	)
}