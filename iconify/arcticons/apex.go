package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.46c.71 0 1.42.66 2.82 2.29l6.51 6.5v-3.68h4v7.69l6.16 6.16l-2.83 2.83L24 10.57L7.33 27.25L4.5 24.42L21.17 7.75c1.42-1.42 2.14-2.23 2.83-2.29Zm13.35 18.46V41.2A1.33 1.33 0 0 1 36 42.54H12a1.33 1.33 0 0 1-1.33-1.34V23.9m22.67-9.65l4 4.01"/>`),
		g.Group(children),
	)
}