package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trunks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 9.2h-3.837L8.619 4.5l-2.457 9.075l1.13 10.882l11.226-3.708L21.835 43.5H24m0-34.3h3.837l11.544-4.7l2.457 9.075l-1.13 10.882l-11.226-3.708L26.164 43.5H24"/><circle cx="19.997" cy="14.573" r=".75" fill="currentColor"/><circle cx="28.003" cy="14.573" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}