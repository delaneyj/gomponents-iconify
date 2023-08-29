package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AsteriskSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="m28.89 20.91l-5-2.91l4.87-2.86a3.11 3.11 0 0 0 1.14-1.08a3 3 0 0 0-4.09-4.15L21 12.76V7a3 3 0 0 0-6 0v5.76l-4.85-2.85a3 3 0 1 0-3 5.18l5 2.91l-4.95 2.86a3.11 3.11 0 0 0-1.14 1.08a3 3 0 0 0 4.09 4.14L15 23.24v5.66a3 3 0 0 0 2 2.94A3 3 0 0 0 21 29v-5.76l4.85 2.85a3 3 0 1 0 3-5.18Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}