package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BillBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2 2h20M8.049 18.53C9.932 20.178 10.873 21 12 21c1.127 0 2.069-.823 3.951-2.47l2-1.748c1.008-.882 1.513-1.322 1.78-1.913c.269-.59.269-1.26.269-2.599V9.702M20 6V2H4v10.27c0 1.34 0 2.009.268 2.6c.175.385.451.707.903 1.13M8.5 13h7m-7-5h7"/>`),
		g.Group(children),
	)
}