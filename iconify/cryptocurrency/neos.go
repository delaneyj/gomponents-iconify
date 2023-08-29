package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Neos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zM10.5 9.358V26h2.429V13.948l5.714 3.433v-3.097zm11 13.284V6h-2.429v12.052l-5.714-3.433v3.097z"/>`),
		g.Group(children),
	)
}