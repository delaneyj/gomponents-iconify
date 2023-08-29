package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectronicPen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="12" height="38" x="35.193" y="5.322" fill="#2F88FF" stroke="#000" stroke-width="4" rx="6" transform="rotate(45 35.193 5.322)"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 23L26 31"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 43L12 37"/><rect width="4" height="4" x="33.268" y="12.34" fill="#fff" rx="2" transform="rotate(30 33.268 12.34)"/></g>`),
		g.Group(children),
	)
}