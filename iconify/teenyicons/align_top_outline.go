package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTopOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15 .5H0m11.5 3h-3v7h3v-7Zm-5 0h-3v11h3v-11Z"/>`),
		g.Group(children),
	)
}