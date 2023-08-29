package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intersex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1024 32q0-14 9-23t23-9h288q26 0 45 19t19 45v288q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V218l-254 255q126 158 126 359q0 221-147.5 384.5T640 1404v132h96q14 0 23 9t9 23v64q0 14-9 23t-23 9h-96v96q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-96h-96q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96v-132q-149-16-270.5-103T55 1077.5T2 786q16-204 160-353.5T509 260q118-14 228 19t198 103l255-254h-134q-14 0-23-9t-9-23V32zM576 1280q185 0 316.5-131.5T1024 832T892.5 515.5T576 384T259.5 515.5T128 832t131.5 316.5T576 1280z"/>`),
		g.Group(children),
	)
}