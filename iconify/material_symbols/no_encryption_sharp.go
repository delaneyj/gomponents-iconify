package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoEncryptionSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 17.15L10.85 8H15V6q0-1.25-.862-2.125T12.025 3q-1.25 0-2.1.863T9.075 6v.225L7.25 4.4q.55-1.525 1.875-2.463t2.9-.937Q14.1 1 15.55 2.463T17 6v2h3v9.15Zm.5 6.15L19.2 22H4V8.1h1.3L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4Zm-9.9-9.85q-.275.275-.425.638t-.15.787q0 .825.588 1.413t1.412.587q.425 0 .788-.15t.637-.425l-2.85-2.85Z"/>`),
		g.Group(children),
	)
}