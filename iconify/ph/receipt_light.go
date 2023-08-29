package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReceiptLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M182 104a6 6 0 0 1-6 6H80a6 6 0 0 1 0-12h96a6 6 0 0 1 6 6Zm-6 26H80a6 6 0 0 0 0 12h96a6 6 0 0 0 0-12Zm54-74v152a6 6 0 0 1-2.85 5.1a5.93 5.93 0 0 1-3.15.9a6 6 0 0 1-2.68-.63L192 198.71l-29.32 14.66a6 6 0 0 1-5.36 0L128 198.71l-29.32 14.66a6 6 0 0 1-5.36 0L64 198.71l-29.32 14.66A6 6 0 0 1 26 208V56a14 14 0 0 1 14-14h176a14 14 0 0 1 14 14Zm-12 0a2 2 0 0 0-2-2H40a2 2 0 0 0-2 2v142.29l23.32-11.66a6 6 0 0 1 5.36 0L96 201.29l29.32-14.66a6 6 0 0 1 5.36 0L160 201.29l29.32-14.66a6 6 0 0 1 5.36 0L218 198.29Z"/>`),
		g.Group(children),
	)
}