package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottomOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15 14.5H0m11.5-3h-3v-7h3v7Zm-5 0h-3V.5h3v11Z"/>`),
		g.Group(children),
	)
}