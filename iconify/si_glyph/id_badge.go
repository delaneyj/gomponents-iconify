package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.969 2.021H2.021C.911 2.021.01 2.916.01 4.016V11c0 1.101.9 1.995 2.011 1.995h.95v-1.042h2.027v1.042h5.961v-1.042h2.051v1.042h.959c1.11 0 2.01-.895 2.01-1.995V4.016c0-1.101-.9-1.995-2.01-1.995zM5.457 3.969c.827 0 1.5.688 1.5 1.534c0 .847-.674 2.466-1.5 2.466c-.827 0-1.5-1.62-1.5-2.466c0-.846.673-1.534 1.5-1.534zm3.537 7.05H1.987s-.12-3.028 1.679-3.028c.374.393.892.77 1.842.77c.951 0 1.404-.381 1.775-.778c2.025-.002 1.711 3.036 1.711 3.036zm5.035-1.066H9.984v-1h4.045v1zm-2.06-1.984h-2v-1h2v1zm2.047-2.015H9.969v-1h4.047v1z"/>`),
		g.Group(children),
	)
}