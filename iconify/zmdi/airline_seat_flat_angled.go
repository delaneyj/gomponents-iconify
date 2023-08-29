package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatFlatAngled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m443 241l-15 40l-264-95l45-121l182 66q34 12 49 44t3 66zM0 195l15-40l405 146l-14 40l-97-34v34H139v-96zm124-41.5q-24 11.5-49 3T38.5 124t-3-49T68 38.5t49-3T153.5 68t3 49t-32.5 36.5z"/>`),
		g.Group(children),
	)
}