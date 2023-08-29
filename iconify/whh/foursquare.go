package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Foursquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M998 327L824 485l58 58q14 14 14 34t-14 34L545 947q-14 14-33.5 14T478 947L163 632q-72-21-117.5-80T0 417q0-92 65.5-158T224 193q54 0 100.5 24t77.5 66l76-76q14-14 33.5-14t33.5 14l137 136L998 5q13-14 22 15q4 11 4 21v226q0 13-7.5 30.5T998 327zm-458-58q-12-12-28.5-12T484 269l-57 56q21 45 21 92q0 37-14 76l54 57l161-172zm249 248L514 766q-7 6-20 1.5T469 754l-12-8l-131-130q-41 21-88 24l246 245q11 12 27.5 12t28.5-12l280-280q12-12 12-28t-12-28z"/>`),
		g.Group(children),
	)
}