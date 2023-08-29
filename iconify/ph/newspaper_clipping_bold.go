package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperClippingBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 36H40a20 20 0 0 0-20 20v160a12 12 0 0 0 17.37 10.73L64 213.42l26.63 13.31a12 12 0 0 0 10.74 0L128 213.42l26.63 13.31a12 12 0 0 0 10.74 0L192 213.42l26.63 13.31A12 12 0 0 0 236 216V56a20 20 0 0 0-20-20Zm-4 160.58l-14.63-7.31a12 12 0 0 0-10.74 0L160 202.58l-26.63-13.31a12 12 0 0 0-10.74 0L96 202.58l-26.63-13.31a12 12 0 0 0-10.74 0L44 196.58V60h168ZM136 108a12 12 0 0 1 12-12h36a12 12 0 0 1 0 24h-36a12 12 0 0 1-12-12Zm0 40a12 12 0 0 1 12-12h36a12 12 0 0 1 0 24h-36a12 12 0 0 1-12-12Zm-64 24h40a12 12 0 0 0 12-12V96a12 12 0 0 0-12-12H72a12 12 0 0 0-12 12v64a12 12 0 0 0 12 12Zm12-64h16v40H84Z"/>`),
		g.Group(children),
	)
}