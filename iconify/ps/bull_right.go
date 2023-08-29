package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BullRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0Q150 0 75 62.5T0 213t75 151t181 63h21q4 0 17 13.5t28.5 29t41.5 29t54 13.5q8 0 10-13t-1.5-32t-8.5-38t-9-32l-4-13q48-30 77.5-74.5T512 213q0-88-75-150.5T256 0zm126 348l-28 17l11 32q14 48 17 62q-26-14-51-45q-2-1-8-7t-8.5-7.5t-8-5.5t-9-5.5t-9-3T277 384h-23q-88 0-150.5-50.5T41 213t62.5-120T254 43t150.5 50T467 213q3 81-85 135z"/>`),
		g.Group(children),
	)
}