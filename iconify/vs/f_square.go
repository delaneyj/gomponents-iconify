package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 1h1120q139 0 237.5 98t98.5 237v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 99T336 1zm339 545h584q31 0 53-21.5t22-52.5v-98q0-31-22-53t-53-22H502q-30 0-52 22t-22 53v1045q0 31 22 52.5t52 21.5h98q31 0 53-21.5t22-52.5v-399h476q31 0 52.5-22t21.5-53v-98q0-31-21.5-53t-52.5-22H675V546z"/>`),
		g.Group(children),
	)
}