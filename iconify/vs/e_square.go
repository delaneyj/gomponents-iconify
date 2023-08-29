package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ESquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M336 0h1120q139 0 237.5 98.5T1792 336v1120q0 139-98.5 237.5T1456 1792H336q-139 0-237.5-98.5T0 1456V336Q0 197 98.5 98.5T336 0zm339 546h584q31 0 53-21.5t22-52.5v-98q0-31-22-53t-53-22H502q-30 0-52 22t-22 53v1044q0 31 22 53t52 22h757q31 0 53-22t22-53v-98q0-31-22-52.5t-53-21.5H677l-2-226h423q31 0 53-22t22-53v-98q0-31-22-53t-53-22H675V546z"/>`),
		g.Group(children),
	)
}