package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdminMedsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.35 16.5q-1 1-2.425 1t-2.425-1q-1-1-1-2.413t1-2.412L11.65 7.5q1-1 2.425-1t2.425 1q1 1 1 2.413t-1 2.412L12.35 16.5ZM8.9 15.075q.575.575 1.188.413t.837-.388l1.375-1.4l-2-2l-1.4 1.375q-.425.425-.425 1t.425 1Zm6.2-6.15q-.575-.575-1.187-.413t-.838.388L11.7 10.3l2 2l1.4-1.375q.425-.425.425-1t-.425-1ZM3 21V3h6.2q.325-.9 1.088-1.45T12 1q.95 0 1.713.55T14.8 3H21v18H3Zm2-2h14V5H5v14Zm7-14.75q.325 0 .537-.213t.213-.537q0-.325-.213-.537T12 2.75q-.325 0-.537.213t-.213.537q0 .325.213.537T12 4.25ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}