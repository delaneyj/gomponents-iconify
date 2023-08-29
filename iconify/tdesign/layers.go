package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1.917l10.1 4.208L12 10.333L1.9 6.125L12 1.917ZM7.1 6.125L12 8.166l4.9-2.041L12 4.083L7.1 6.125ZM2 9.98l10 4.308L22 9.98v2.181l-9.604 4.134l-.396.17l-.395-.17L2 12.162v-2.18Zm0 6l10 4.308l10-4.308v2.181l-9.603 4.134l-.397.17l-.395-.17L2 18.162v-2.18Z"/>`),
		g.Group(children),
	)
}