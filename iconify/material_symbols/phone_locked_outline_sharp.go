package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneLockedOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 10V5h1V4q0-.825.588-1.413T18 2q.825 0 1.413.588T20 4v1h1v5h-6Zm2-5h2V4q0-.425-.288-.713T18 3q-.425 0-.713.288T17 4v1Zm2.95 16q-3.125 0-6.188-1.35T8.2 15.8q-2.5-2.5-3.85-5.55T3 4.05V3h5.9l.925 5.025l-2.85 2.875q.55.975 1.225 1.85t1.45 1.625q.725.725 1.588 1.388T13.1 17l2.9-2.9l5 1.025V21h-1.05ZM6.025 9l1.65-1.65L7.25 5H5.025q.125 1.125.375 2.113T6.025 9Zm8.95 8.95q1 .425 2.013.675T19 18.95v-2.2l-2.35-.475l-1.675 1.675ZM6.025 9Zm8.95 8.95Z"/>`),
		g.Group(children),
	)
}