package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletAndroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M320 0q27 0 45.5 18.5T384 64v384q0 27-18.5 45.5T320 512H64q-27 0-45.5-18.5T0 448V64q0-27 18.5-45.5T64 0h256zm-85 469v-21h-86v21h86zm112-64V64H37v341h310z"/>`),
		g.Group(children),
	)
}