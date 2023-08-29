package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Interactions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 18h-4v-4h-2v14h6a2.003 2.003 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2zm-4 8v-6h4v6zM20 6.076l-.744-1.857L16 5.522V2h-2v3.523L10.744 4.22L10 6.077l3.417 1.367L10.9 10.8l1.6 1.2L15 8.667L17.5 12l1.6-1.2l-2.517-3.357L20 6.076zM10 18H5v2h5v2H6a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h6v-8a2.002 2.002 0 0 0-2-2zm0 8H6v-2h4z"/>`),
		g.Group(children),
	)
}