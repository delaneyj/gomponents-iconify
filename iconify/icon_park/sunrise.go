package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24H7"/><path d="M10 10L12 12"/><path d="M24 4V7"/><path d="M14 24C14 18.4776 18.4776 14 24 14C29.5224 14 34 18.4776 34 24C34 27.3674 32.3357 30.3458 29.785 32.1578"/><path d="M38 10L36 12"/><path d="M44 24L41 24"/><path d="M37.9814 37.982L36.3614 36.362"/><path d="M23.4999 28C20.4999 28 14 28.2 14 31C14 33.8 18.6058 33.7908 20.9998 34C23 34.1747 26.4624 35.6879 25.9999 38C24.9998 43 8.99982 42 4.99994 42"/></g>`),
		g.Group(children),
	)
}