package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivacySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.85 7.025q.05 0 .075-.013Q9.95 7 10 7q.825 0 1.413.587Q12 8.175 12 9v.175ZM22 17.5l-4-4v1.675l-2-2V6H8.825l-2-2H16q.825 0 1.413.588Q18 5.175 18 6v4.5l4-4ZM16 20H4q-.825 0-1.412-.587Q2 18.825 2 18V6q0-.825.588-1.412Q3.175 4 4 4l2 2H4v12h5.3v-1.05q-1.875-.275-3.087-1.675Q5 13.875 5 12h1.425q0 1.5 1.038 2.537Q8.5 15.575 10 15.575q.825 0 1.562-.363q.738-.362 1.238-1.012l1.025 1.025q-.6.725-1.4 1.162q-.8.438-1.725.563V18H16v-2l2 2q0 .825-.587 1.413Q16.825 20 16 20Zm5.95 1.95l-1.4 1.4l-9.625-9.625q-.225.125-.45.2Q10.25 14 10 14q-.825 0-1.412-.588Q8 12.825 8 12v-1.2L.65 3.45l1.4-1.4Z"/>`),
		g.Group(children),
	)
}