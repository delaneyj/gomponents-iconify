package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quakedroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.338 30.059A14.894 14.894 0 0 1 9.145 15.264A13.92 13.92 0 0 1 14.057 4.5a14.37 14.37 0 0 0-2.882 8.458a12.865 12.865 0 0 0 11.12 12.745l.074.009v-3.827c0-2.231-2.47-2.47-2.47-2.47h8.275s-2.469.239-2.469 2.47v3.828h0a12.865 12.865 0 0 0 11.12-12.745a14.37 14.37 0 0 0-2.882-8.458a13.92 13.92 0 0 1 4.912 10.764A14.894 14.894 0 0 1 25.662 30.07l.043-.007v4.586L24 43.5l-1.631-8.852V30.09Z"/>`),
		g.Group(children),
	)
}