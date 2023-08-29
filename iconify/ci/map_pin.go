package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 9.923c0 4.852 4.244 8.864 6.123 10.402c.27.22.405.332.606.388c.156.044.386.044.542 0c.201-.056.336-.167.606-.388C14.756 18.787 19 14.775 19 9.923a6.885 6.885 0 0 0-2.05-4.895A7.04 7.04 0 0 0 12 3a7.04 7.04 0 0 0-4.95 2.028A6.884 6.884 0 0 0 5 9.923Z"/><path d="M10 9a2 2 0 1 0 4 0a2 2 0 0 0-4 0Z"/></g>`),
		g.Group(children),
	)
}