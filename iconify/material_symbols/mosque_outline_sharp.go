package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MosqueOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21V8.725Q.55 8.45.275 8.012T0 7q0-.575.6-1.4T2 4q.8.775 1.4 1.6T4 7q0 .575-.275 1.012T3 8.725V13h2V9.05h1.55q-.275-.425-.413-.925T6 7.1q0-1 .475-1.85t1.275-1.4L12 1l4.25 2.85q.8.55 1.275 1.4T18 7.1q0 .525-.138 1.025t-.412.925H19V13h2V8.725q-.45-.275-.725-.713T20 7q0-.575.6-1.4T22 4q.8.775 1.4 1.6T24 7q0 .575-.275 1.012T23 8.725V21H13v-5h-2v5H1ZM9.9 9h4.2q.8 0 1.35-.55T16 7.1q0-.5-.225-.913T15.15 5.5L12 3.4L8.85 5.5q-.4.275-.625.688T8 7.1q0 .8.55 1.35T9.9 9ZM3 19h6v-5h6v5h6v-4h-4v-4H7v4H3v4Zm9-8Zm0-2Zm0 .05Z"/>`),
		g.Group(children),
	)
}