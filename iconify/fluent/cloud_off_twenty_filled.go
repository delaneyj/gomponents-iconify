package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudOffTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.854 2.146a.5.5 0 1 0-.708.708l3.67 3.668a5.326 5.326 0 0 0-.463 1.724h-.07C3.468 8.246 2 9.758 2 11.623C2 13.488 3.47 15 5.282 15h9.01l2.854 2.854a.5.5 0 0 0 .708-.708l-15-15ZM18 11.623a3.4 3.4 0 0 1-1.452 2.804l-9.49-9.49C7.808 4.353 8.792 4 10 4c2.817 0 4.415 1.923 4.647 4.246h.07c1.814 0 3.283 1.512 3.283 3.377Z"/>`),
		g.Group(children),
	)
}