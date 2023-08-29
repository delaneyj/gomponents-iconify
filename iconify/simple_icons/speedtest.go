package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speedtest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.628 16.186l-2.047-2.14l6.791-5.953l1.21 1.302zm8.837 6.047c2.14-2.14 3.535-5.117 3.535-8.466c0-6.604-5.395-12-12-12s-12 5.396-12 12c0 3.35 1.302 6.326 3.535 8.466l1.674-1.675c-1.767-1.767-2.79-4.093-2.79-6.79A9.568 9.568 0 0 1 12 4.185a9.568 9.568 0 0 1 9.581 9.581c0 2.605-1.116 5.024-2.79 6.791Z"/>`),
		g.Group(children),
	)
}