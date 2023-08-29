package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zM11.89 6L9 7.284v2.897l7.693 3.737L9 17.728v2.856l11.114 5.351l2.845-1.281v-8.951l-2.845 1.295v6.085l-7.548-3.67l7.548-3.897l2.845-1.281l-.013-2.86z"/><path fill="currentColor" fill-opacity=".65" d="m9 7.281l11.114 5.383l2.845-1.282L11.891 6z"/><path fill="currentColor" fill-opacity=".3" d="m20.114 12.651l2.845-1.281v2.865l-2.845 1.281zm0 13.284v-8.937l2.845-1.295v8.951z"/>`),
		g.Group(children),
	)
}