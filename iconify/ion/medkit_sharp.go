package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedkitSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" d="M168 72h176v24H168z"/><path fill="currentColor" d="M484 96H384V40a8 8 0 0 0-8-8H136a8 8 0 0 0-8 8v56H28a12 12 0 0 0-12 12v360a12 12 0 0 0 12 12h456a12 12 0 0 0 12-12V108a12 12 0 0 0-12-12ZM168 72h176v24H168Zm184 238h-74v74h-44v-74h-74v-44h74v-74h44v74h74Z"/>`),
		g.Group(children),
	)
}