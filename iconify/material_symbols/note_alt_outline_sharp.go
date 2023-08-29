package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteAltOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h2.1l6-5.95l-2.15-2.15L7 14.85V17Zm8.8-6.65l1.4-1.45l-2.1-2.1l-1.45 1.4l2.15 2.15ZM3 21V3h6.2q.325-.9 1.088-1.45T12 1q.95 0 1.713.55T14.8 3H21v18H3Zm2-2h14V5H5v14Zm7-14.75q.325 0 .537-.213t.213-.537q0-.325-.213-.537T12 2.75q-.325 0-.537.213t-.213.537q0 .325.213.537T12 4.25ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}