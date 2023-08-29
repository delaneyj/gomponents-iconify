package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Texture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.4 21q-.475-.1-.888-.513T3 19.6L19.6 3q.525.125.9.513t.525.887L4.4 21ZM3 14.7v-2.8L11.9 3h2.8L3 14.7ZM3 7V5q0-.825.588-1.413T5 3h2L3 7Zm14 14l4-4v2q0 .825-.588 1.413T19 21h-2Zm-7.7 0L21 9.3v2.8L12.1 21H9.3Z"/>`),
		g.Group(children),
	)
}