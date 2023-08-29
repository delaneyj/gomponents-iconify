package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Earthquake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16.612 2.214a1.01 1.01 0 0 0-1.242 0L1 13.419l1.243 1.572L4 13.621V26a2.004 2.004 0 0 0 2 2h20a2.004 2.004 0 0 0 2-2V13.63L29.757 15L31 13.428ZM6 12.062l9-7.012v7.364l3.458 3.458l-6.734 4.81L14.382 26H6ZM26 26h-9.382l-2.342-4.683l7.266-5.19L17 11.587V5.052l9 7.02Z"/>`),
		g.Group(children),
	)
}