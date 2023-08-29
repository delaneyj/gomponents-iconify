package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="m1100 1653l382-382h-260V442q39 70 89 143t91 126t110 131t106 117t122 127.5t116 121.5l-200 200h544V808l-200 200q-70-60-157.5-128T1694 766.5t-130.5-103T1447 559t-93-108.5t-72.5-124.5t-43-143.5T1222 8H978q0 94-16.5 177T919 333.5t-73.5 133t-92 117.5t-117 114t-130 110t-149 120T200 1058L0 858v550h550l-200-200q21-22 99-104t104.5-110.5t91-99t97-111.5t81-103.5t85-119.5T978 442v829H718z"/>`),
		g.Group(children),
	)
}