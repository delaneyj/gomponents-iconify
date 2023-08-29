package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneGavel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21h12v2H1v-2zM5.24 8.07l2.83-2.83l14.14 14.14l-2.83 2.83L5.24 8.07zM12.32 1l5.66 5.66l-2.83 2.83l-5.66-5.66L12.32 1zM3.83 9.48l5.66 5.66l-2.83 2.83L1 12.31l2.83-2.83z"/>`),
		g.Group(children),
	)
}