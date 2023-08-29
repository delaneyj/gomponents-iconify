package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurlingStone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14.686 9.686c-.165.164-.186.302-.186.314v2H21a2 2 0 0 1 2 2v2a7 7 0 1 1 0 14H9a7 7 0 1 1 0-14v-2a2 2 0 0 1 2-2h.5v-2c0-.988.479-1.85 1.064-2.436C13.15 6.98 14.012 6.5 15 6.5h6a1.5 1.5 0 0 1 0 3h-6c-.012 0-.15.021-.314.186ZM4.999 20a4.973 4.973 0 0 0-1 3A5 5 0 0 0 9 28h14a5 5 0 0 0 4-8H5Z"/>`),
		g.Group(children),
	)
}