package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitalYuan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.654 29.924A21.499 21.499 0 1 1 45.486 24M8.287 31.807h31.478M45.5 24.03H2.502m21.715-4.238v25.59M35.65 5.936L24.216 19.792M12.429 5.879l11.788 13.913"/>`),
		g.Group(children),
	)
}