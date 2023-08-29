package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareStopTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M23.75 5a2.25 2.25 0 0 1 2.245 2.095L26 7.25V20.75a2.25 2.25 0 0 1-2.096 2.245L23.75 23H4.25a2.25 2.25 0 0 1-2.245-2.096L2 20.75V7.25a2.25 2.25 0 0 1 2.096-2.245l.154-.006h19.5zm0 1.5H4.25a.75.75 0 0 0-.743.648l-.007.101V20.75c0 .38.282.693.648.743l.102.007h19.5a.75.75 0 0 0 .743-.648l.007-.102V7.25a.75.75 0 0 0-.648-.744L23.75 6.5zm-13.533 3.717a.75.75 0 0 1 .977-.072l.084.072l2.72 2.72l2.725-2.716a.75.75 0 0 1 1.132.978l-.073.084L15.059 14l2.724 2.723a.75.75 0 0 1-.977 1.133l-.084-.072l-2.724-2.724l-2.72 2.723a.75.75 0 0 1-1.134-.975l.073-.085l2.72-2.724l-2.72-2.72a.75.75 0 0 1 0-1.06z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}