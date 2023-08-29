package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HowToVote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20v-4.55l2.75-3.125l1.425 1.425l-2 2.25h13.65l-1.95-2.2l1.425-1.425L21 15.45V20q0 .825-.588 1.413T19 22H5Zm5.6-7.6l-3.475-3.525Q6.55 10.3 6.55 9.45t.575-1.425l4.9-4.9q.575-.575 1.425-.575t1.425.575L18.4 6.6q.575.575.588 1.412t-.563 1.413l-5 5q-.575.575-1.413.563T10.6 14.4ZM17 8l-3.55-3.5L8.5 9.45l3.55 3.5L17 8Z"/>`),
		g.Group(children),
	)
}