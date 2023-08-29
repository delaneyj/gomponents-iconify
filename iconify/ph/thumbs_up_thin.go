package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUpThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M231 82.76A20 20 0 0 0 216 76h-60V56a36 36 0 0 0-36-36a4 4 0 0 0-3.58 2.21L77.53 100H32a12 12 0 0 0-12 12v88a12 12 0 0 0 12 12h172a20 20 0 0 0 19.85-17.52l12-96A20 20 0 0 0 231 82.76ZM76 204H32a4 4 0 0 1-4-4v-88a4 4 0 0 1 4-4h44ZM227.91 97.49l-12 96A12 12 0 0 1 204 204H84v-99.06l38.42-76.84A28 28 0 0 1 148 56v24a4 4 0 0 0 4 4h64a12 12 0 0 1 11.91 13.49Z"/>`),
		g.Group(children),
	)
}