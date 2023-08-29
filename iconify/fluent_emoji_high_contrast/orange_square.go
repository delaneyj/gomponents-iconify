package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrangeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 1h24a3 3 0 0 1 3 3v24a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V4a3 3 0 0 1 3-3ZM3 4v1.645L5.645 3H4a1 1 0 0 0-1 1Zm0 3.059v3.586L10.645 3H7.059L3 7.059ZM12.059 3L3 12.059v3.586L15.645 3h-3.586Zm5 0L3 17.059v3.586L20.645 3h-3.586Zm5 0L3 22.059v3.586L25.645 3h-3.586Zm5 0L3 27.059V28a1 1 0 0 0 1 1h.645L29 4.645V4a1 1 0 0 0-1-1h-.941ZM29 6.059L6.059 29h3.586L29 9.645V6.059Zm0 5L11.059 29h3.586L29 14.645v-3.586Zm0 5L16.059 29h3.586L29 19.645v-3.586Zm0 5L21.059 29h3.586L29 24.645v-3.586Zm0 5L26.059 29H28a1 1 0 0 0 1-1v-1.941Z"/>`),
		g.Group(children),
	)
}