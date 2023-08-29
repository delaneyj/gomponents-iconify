package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bestiary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.75 7.48a17.5 17.5 0 1 0-19.5 29.06v7h4.09v-6.26H22v6.23h4v-6.23h3.63v6.23h4.09v-7a17.5 17.5 0 0 0 0-29.06ZM16.31 30.16a3.92 3.92 0 1 1 3.91-3.92a3.93 3.93 0 0 1-3.91 3.92Zm4.85 3.41L24 27.83l2.84 5.74Zm10.53-3.41a3.92 3.92 0 1 1 3.92-3.92a3.92 3.92 0 0 1-3.92 3.92Z"/>`),
		g.Group(children),
	)
}