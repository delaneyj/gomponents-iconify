package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Store(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<path fill="currentColor" d="M25.562 2.966L0 21h6v22h4V28h8v15h27V21h5L25.562 2.966zM42 37H23v-9h19v9zm0-12H9v-8h33v8zm-2-6v4H11v-4h29m1-1H10v6h31v-6z"/>`),
		g.Group(children),
	)
}