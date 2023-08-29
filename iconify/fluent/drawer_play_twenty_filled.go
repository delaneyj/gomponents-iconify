package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrawerPlayTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8.207 4H6a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V9.743a5.507 5.507 0 0 1-1 .657v.6h-3.5a.5.5 0 0 0-.5.5a2 2 0 1 1-4 0a.5.5 0 0 0-.5-.5H4V8h4.6a5.463 5.463 0 0 1-.393-1H4a2 2 0 0 1 2-2h2.022a5.48 5.48 0 0 1 .185-1Zm5.293 6a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm-.757-6.587l2.97 1.65a.5.5 0 0 1 0 .874l-2.97 1.65A.5.5 0 0 1 12 7.15v-3.3a.5.5 0 0 1 .743-.437Z"/>`),
		g.Group(children),
	)
}