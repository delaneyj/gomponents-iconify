package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoEncryptionOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 17.15l-2-2V10h-5.15l-2-2H15V6q0-1.25-.862-2.125T12.025 3q-1.25 0-2.1.863T9.075 6v.225L7.25 4.4q.55-1.525 1.875-2.463t2.9-.937Q14.1 1 15.55 2.463T17 6v2h3v9.15Zm.5 6.15L19.2 22H4V8.1h1.3L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM17.15 20l-3.675-3.725q-.275.275-.637.425t-.788.15q-.825 0-1.412-.588t-.588-1.412q0-.425.15-.787t.425-.638L7.15 10H6v10h11.15Zm-5-5Zm3.275-2.425Z"/>`),
		g.Group(children),
	)
}