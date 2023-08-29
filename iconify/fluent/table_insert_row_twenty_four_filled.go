package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableInsertRowTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M20.25 4.5H3.75a.75.75 0 0 1 0-1.5h16.5a.75.75 0 0 1 0 1.5ZM8 8H5.25C4.007 8 3 8.895 3 10v4c0 1.105 1.007 2 2.25 2H8V8Zm1.5 8h5V8h-5v8Zm9.25 0H16V8h2.75c1.243 0 2.25.895 2.25 2v4c0 1.105-1.007 2-2.25 2Zm-15 5h16.5a.75.75 0 0 0 0-1.5H3.75a.75.75 0 0 0 0 1.5Z"/>`),
		g.Group(children),
	)
}