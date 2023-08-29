package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fruiter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M19 32.0003C15 32 6 32.1309 6 22.941C6 21.192 6.59395 17.6967 9.8653 15.6125C10.5277 15.1905 10.9894 14.4933 11.1004 13.7159C11.5991 10.2257 14.1089 4 22 4C24.2925 4 27.9884 4.29995 30.3835 7.93C30.8001 8.56136 31.4798 8.98871 32.2221 9.13376C35.7901 9.83088 42.0582 13.2757 42 20C42.0591 22.1704 41.5574 25.3457 37.494 27.3697C36.6482 27.791 36.1153 28.6832 35.9478 29.6132C35.497 32.117 33.2767 35.3748 27 36"/><path fill="#2F88FF" d="M16 44C23.0933 31.0694 18.9556 25.3469 16 24L30 22C23.28 31.3388 26.5778 40.5578 29.0667 44H16Z"/></g>`),
		g.Group(children),
	)
}