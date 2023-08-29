package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChiliAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.25 7.5l-1.52-.87C8.26 5.7 9.03 5 9.94 4.69C9.8 4.29 9.44 4 9 4V2c1.54 0 2.79 1.16 2.97 2.65c.94.29 1.75 1.01 2.3 1.98l-1.52.87l-1.75-1l-1.75 1m4.75.78L12.75 9L11 8L9.25 9L8 8.28c-.6.35-1 .99-1 1.72v1c0 9 8 11 8 11V10c0-.73-.4-1.37-1-1.72M17 7v6h2V7h-2m0 10h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}