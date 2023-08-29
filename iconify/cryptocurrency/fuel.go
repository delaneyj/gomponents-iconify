package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fuel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm9.532-22h-6.495l-1.926 4.83h-6.074v-2.487h6.074L18 10H8.802a.43.43 0 0 0-.432.427v4.402H6.89L6 17.171h2.37v4.397c0 .238.196.432.438.432h4.451l.889-2.341h-3.111V17.17h5.037L14.148 22h2.963l2.074-4.83h6.38c.24 0 .435-.192.435-.429v-6.28a.465.465 0 0 0-.468-.461zm-2.05 4.46c0 .204-.167.37-.374.37h-3.034l1.037-2.486h1.995a.373.373 0 0 1 .377.368z"/>`),
		g.Group(children),
	)
}