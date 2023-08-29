package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.05 22v-2.05q.85-.125 1.663-.45t1.537-.85l1.4 1.45q-1.05.8-2.2 1.287t-2.4.613Zm-2 0q-3.45-.45-5.725-2.988T3.05 13.05q0-1.875.713-3.513t1.925-2.85Q6.9 5.476 8.537 4.764t3.512-.713h.15L10.65 2.5l1.4-1.45l4 4l-4 4l-1.4-1.4l1.6-1.6h-.2q-2.925 0-4.963 2.038T5.05 13.05q0 2.6 1.7 4.563t4.3 2.337V22Zm8.05-3.35l-1.45-1.4q.525-.725.85-1.538t.45-1.662H21q-.125 1.25-.612 2.4t-1.288 2.2Zm1.9-6.6h-2.05q-.125-.85-.45-1.663t-.85-1.537l1.45-1.4q.8.975 1.275 2.15T21 12.05Z"/>`),
		g.Group(children),
	)
}