package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ButtonPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.309 3L9.032 4.125L3.63 3C2.178 3 1.001 4.329 1.001 5.967v5.021c0 1.638 1.177 2.967 2.629 2.967l5.402-1.018l5.277 1.018c1.452 0 2.629-1.329 2.629-2.967V5.967C16.938 4.329 15.761 3 14.309 3zm-3.48 6.117l-2.505 1.917a1.167 1.167 0 0 1-1.418-.004l.029-5.104c.419-.319.945-.222 1.368.101l2.521 1.93c.422.321.424.841.005 1.16z"/>`),
		g.Group(children),
	)
}