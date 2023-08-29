package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xuc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm11-16c0-6.075-4.925-11-11-11S5 9.924 5 16c0 6.075 4.925 11 11 11s11-4.925 11-11zm-3.77.346l-12.318.017a5.094 5.094 0 0 0 5.077 4.726a5.08 5.08 0 0 0 4.48-2.686h2.322a7.23 7.23 0 0 1-5.44 4.681v1.446l-2.24.698V23.16c-3.574-.435-6.345-3.474-6.345-7.165c0-3.605 2.647-6.587 6.103-7.125V7.467l2.239-.697v2.1c3.457.538 6.122 3.742 6.122 7.35c0 .126.007 0 0 .126zM11.25 14.13h9.477a5.095 5.095 0 0 0-4.738-3.229a5.093 5.093 0 0 0-4.738 3.23z"/>`),
		g.Group(children),
	)
}