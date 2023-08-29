package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playground(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 23.5c3-5 3-14.5 3-20.578V.5h17v2.422c0 6.078 0 15.578 3 20.578M9.5.5v15m5-15v15m-8 0v3h11v-3h-11Z"/>`),
		g.Group(children),
	)
}