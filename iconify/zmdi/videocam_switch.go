package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocamSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m341 139l86-86v278l-86-86v75q0 9-6 15t-15 6H21q-8 0-14.5-6T0 320V64q0-9 6.5-15T21 43h299q9 0 15 6t6 15v75zM235 267l74-75l-74-75v54H107v-54l-75 75l75 75v-54h128v54z"/>`),
		g.Group(children),
	)
}