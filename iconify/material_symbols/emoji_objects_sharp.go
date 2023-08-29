package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiObjectsSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-.65 0-1.175-.313T10 20.85H8V15.3q-1.475-.975-2.363-2.575T4.75 9.25q0-3.025 2.113-5.138T12 2q3.025 0 5.138 2.113T19.25 9.25q0 1.925-.888 3.5T16 15.3v5.55h-2q-.3.525-.825.838T12 22Zm-2-3.15h4v-.9h-4v.9Zm0-1.9h4V16h-4v.95ZM11.25 14h1.5v-2.7l2.2-2.2l-1.05-1.05l-1.9 1.9l-1.9-1.9L9.05 9.1l2.2 2.2V14Z"/>`),
		g.Group(children),
	)
}