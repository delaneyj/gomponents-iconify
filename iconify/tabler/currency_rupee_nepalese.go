package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrencyRupeeNepalese(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 5H4h3a4 4 0 1 1 0 8H4l6 6m11-2l-4.586-4.414a2 2 0 0 0-2.828 2.828l.707.707"/>`),
		g.Group(children),
	)
}