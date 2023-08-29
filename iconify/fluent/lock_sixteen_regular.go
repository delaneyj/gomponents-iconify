package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M7.836 1.505L8 1.5a2.5 2.5 0 0 1 2.495 2.336L10.5 4v1h1A1.5 1.5 0 0 1 13 6.5v6a1.5 1.5 0 0 1-1.5 1.5h-7A1.5 1.5 0 0 1 3 12.5v-6A1.5 1.5 0 0 1 4.5 5h1V4a2.5 2.5 0 0 1 2.336-2.495zM11.5 6h-7a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h7a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5zM8 8.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm.144-5.993L8 2.5a1.5 1.5 0 0 0-1.493 1.356L6.5 4v1h3V4a1.5 1.5 0 0 0-1.356-1.493z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}