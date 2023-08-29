package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoPad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.44 2.065H6.56a2.507 2.507 0 0 0-2.5 2.5v14.87a2.507 2.507 0 0 0 2.5 2.5h10.88a2.5 2.5 0 0 0 2.5-2.5V4.565a2.5 2.5 0 0 0-2.5-2.5Zm1.5 17.37a1.5 1.5 0 0 1-1.5 1.5H6.56a1.5 1.5 0 0 1-1.5-1.5V6.505h13.88Z"/><path fill="currentColor" d="M7.549 9.506a.5.5 0 0 1 0-1h8.909a.5.5 0 0 1 0 1Zm0 3a.5.5 0 0 1 0-1h6.5a.5.5 0 0 1 0 1Zm.017 5.868a.5.5 0 1 1 0-1h3.251a.5.5 0 0 1 0 1Z"/>`),
		g.Group(children),
	)
}