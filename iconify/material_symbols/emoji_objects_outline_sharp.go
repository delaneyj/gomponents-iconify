package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiObjectsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.65 0-1.175-.313T10 20.85H8V15.3q-1.475-.975-2.363-2.575T4.75 9.25q0-3.025 2.113-5.138T12 2q3.025 0 5.138 2.113T19.25 9.25q0 1.925-.888 3.5T16 15.3v5.55h-2q-.3.525-.825.838T12 22Zm-2-3.15h4v-.9h-4v.9Zm0-1.9h4V16h-4v.95ZM9.8 14h1.45v-2.7l-2.2-2.2l1.05-1.05l1.9 1.9l1.9-1.9l1.05 1.05l-2.2 2.2V14h1.45q1.35-.65 2.2-1.913t.85-2.837q0-2.2-1.525-3.725T12 4Q9.8 4 8.275 5.525T6.75 9.25q0 1.575.85 2.838T9.8 14ZM12 9.95ZM12 9Z"/>`),
		g.Group(children),
	)
}