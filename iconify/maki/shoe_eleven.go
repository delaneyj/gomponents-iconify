package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShoeEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7.5 5L5.805 3.983L4.8 2.676a.446.446 0 0 0-.627-.083a.442.442 0 0 0-.164.392h-.008V4H2v-.5a.5.5 0 1 0-1 0V6h3.5C5.5 6 6 7 7 7h3l.004-.5C10.004 5.946 7.5 5 7.5 5z" fill="currentColor"/><path d="M5.527 7.584C5.117 7.31 4.651 7 4 7H1v1.47c0 .293.237.53.53.53h1.94A.53.53 0 0 0 4 8.47V8c.349 0 .638.192.973.416c.41.274.876.584 1.527.584h3a.5.5 0 0 0 .5-.5V8H6.5c-.349 0-.638-.193-.973-.416z" fill="currentColor"/>`),
		g.Group(children),
	)
}