package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M14.5 0v15m-3-11.5v3h-7v-3h7Zm0 5v3H.5v-3h11Z"/>`),
		g.Group(children),
	)
}