package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gif(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M64 136v224a32.036 32.036 0 0 0 32 32h96V232h-64v32h32v96H96V136h96v-32H96a32.036 32.036 0 0 0-32 32Zm176-32h32v288h-32zm208 32v-32H328v288h32V264h72v-32h-72v-96h88z"/>`),
		g.Group(children),
	)
}