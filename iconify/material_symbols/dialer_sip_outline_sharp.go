package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialerSipOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.95 21q-3.125 0-6.187-1.35T8.2 15.8q-2.5-2.5-3.85-5.55T3 4.05V3h5.9l.925 5.025l-2.85 2.875q.55.975 1.225 1.85t1.45 1.625q.725.725 1.588 1.388T13.1 17l2.9-2.9l5 1.025V21h-1.05ZM6.025 9l1.65-1.65L7.25 5H5.025q.125 1.125.375 2.113T6.025 9Zm8.95 8.95q1 .425 2.013.675T19 18.95v-2.2l-2.35-.475l-1.675 1.675ZM6.025 9Zm8.95 8.95ZM12 8V7h2V6h-2V3h3v1h-2v1h2v3h-3Zm4 0V3h1v5h-1Zm2 0V3h3v3h-2v2h-1Zm1-3h1V4h-1v1Z"/>`),
		g.Group(children),
	)
}