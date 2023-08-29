package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenBrush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.999 22h3.623a3 3 0 0 0 2.12-.878l14.79-14.789l-4.866-4.865L2.878 16.256a3 3 0 0 0-.879 2.122V22Zm2-2v-1.622a1 1 0 0 1 .293-.707l2.158-2.159l2.037 2.037l-2.158 2.158a1 1 0 0 1-.707.293H3.999Zm5.902-3.865l-2.037-2.037l9.802-9.801l2.037 2.036L9.9 16.135Z"/>`),
		g.Group(children),
	)
}