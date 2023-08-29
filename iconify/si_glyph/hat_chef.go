package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HatChef(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 15h10v1H4zm9.449-11.688A4.243 4.243 0 0 0 9.322.086a4.244 4.244 0 0 0-4.166 3.39a3.28 3.28 0 0 0-.833-.119a3.31 3.31 0 0 0-3.312 3.312c0 1.829 1.187 3.312 3.017 3.312c.019 0 .036-.005.056-.005v3.981h9.875v-3.384c1.883-.14 3.011-1.709 3.011-3.627a3.642 3.642 0 0 0-3.521-3.634zM8.9 1.795c-.842.544-1.383 1.401-1.383 2.368H6.216c0-1.398 1.133-2.571 2.664-2.9c.277-.058.452.252.02.532z"/>`),
		g.Group(children),
	)
}