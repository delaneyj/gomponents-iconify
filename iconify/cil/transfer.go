package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m356.687 228.687l22.626 22.626L494.627 136L379.313 20.687l-22.626 22.626L433.372 120H16v32h417.372l-76.685 76.687zM496 360H78.628l76.685-76.687l-22.626-22.626L17.373 376l115.314 115.313l22.626-22.626L78.628 392H496v-32z"/>`),
		g.Group(children),
	)
}