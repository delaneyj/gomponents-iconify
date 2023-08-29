package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForNorthKorea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m16.362 29.795l-5.995.018l4.842 3.721l-1.839 6.003l4.863-3.691l4.863 3.691l-1.842-6.003l4.846-3.721l-5.998-.018l-1.869-5.99z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2M9.551 48.717a28.193 28.193 0 0 1-2.45-3.934H56.9a28.251 28.251 0 0 1-2.45 3.934H9.551m44.897-33.434a28.041 28.041 0 0 1 2.45 3.934H7.102a28.1 28.1 0 0 1 2.45-3.934h44.896M29.05 32c0 5.974-4.844 10.817-10.816 10.817c-5.975 0-10.816-4.843-10.816-10.817s4.842-10.817 10.816-10.817c5.972 0 10.816 4.842 10.816 10.817"/>`),
		g.Group(children),
	)
}