package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mercury(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M830 316q145 72 233.5 210.5T1152 832q0 221-147.5 384.5T640 1404v132h96q14 0 23 9t9 23v64q0 14-9 23t-23 9h-96v96q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23v-96h-96q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96v-132q-217-24-364.5-187.5T0 832q0-167 88.5-305.5T322 316Q157 220 94 43q-6-16 3.5-29.5T124 0h69q21 0 29 20q44 106 140 171t214 65t214-65T930 20q8-20 37-20h61q17 0 26.5 13.5T1058 43q-63 177-228 273zm-254 964q185 0 316.5-131.5T1024 832T892.5 515.5T576 384T259.5 515.5T128 832t131.5 316.5T576 1280z"/>`),
		g.Group(children),
	)
}