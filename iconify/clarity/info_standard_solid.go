package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoStandardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M18 2.1a16 16 0 1 0 16 16a16 16 0 0 0-16-16Zm-.1 5.28a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm3.6 21.25h-7a1.4 1.4 0 1 1 0-2.8h2.1v-9.2H15a1.4 1.4 0 1 1 0-2.8h4.4v12h2.1a1.4 1.4 0 1 1 0 2.8Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}