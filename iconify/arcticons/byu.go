package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Byu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.355 7.708c-2.53 1.997-5.715 30.295 1.283 30.302c8.913.01 14.973-33.51 14.973-33.51c-1.782 1.782-7.557 34.865-1.782 39"/>`),
		g.Group(children),
	)
}