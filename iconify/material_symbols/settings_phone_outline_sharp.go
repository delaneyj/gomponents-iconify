package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsPhoneOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 11q-.425 0-.713-.288T11 10q0-.425.288-.713T12 9q.425 0 .713.288T13 10q0 .425-.288.713T12 11Zm4 0q-.425 0-.713-.288T15 10q0-.425.288-.713T16 9q.425 0 .713.288T17 10q0 .425-.288.713T16 11Zm4 0q-.425 0-.713-.288T19 10q0-.425.288-.713T20 9q.425 0 .713.288T21 10q0 .425-.288.713T20 11Zm-.05 10q-3.125 0-6.188-1.35T8.2 15.8q-2.5-2.5-3.85-5.55T3 4.05V3h5.9l.925 5.025l-2.85 2.875q.55.975 1.225 1.85t1.45 1.625q.725.725 1.588 1.388T13.1 17l2.9-2.9l5 1.025V21h-1.05ZM6.025 9l1.65-1.65L7.25 5H5.025q.125 1.125.375 2.113T6.025 9Zm8.95 8.95q1 .425 2.013.675T19 18.95v-2.2l-2.35-.475l-1.675 1.675ZM6.025 9Zm8.95 8.95Z"/>`),
		g.Group(children),
	)
}