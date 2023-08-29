package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaylistAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 152h208v32H256zm-80 104h288v32H176zm0 104h288v32H176zm16-208h-64V88H96v64H32v32h64v64h32v-64h64v-32z"/>`),
		g.Group(children),
	)
}