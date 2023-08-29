package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserFocusLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M222 48v28a6 6 0 0 1-12 0V48a2 2 0 0 0-2-2h-28a6 6 0 0 1 0-12h28a14 14 0 0 1 14 14Zm-6 126a6 6 0 0 0-6 6v28a2 2 0 0 1-2 2h-28a6 6 0 0 0 0 12h28a14 14 0 0 0 14-14v-28a6 6 0 0 0-6-6ZM76 210H48a2 2 0 0 1-2-2v-28a6 6 0 0 0-12 0v28a14 14 0 0 0 14 14h28a6 6 0 0 0 0-12ZM40 82a6 6 0 0 0 6-6V48a2 2 0 0 1 2-2h28a6 6 0 0 0 0-12H48a14 14 0 0 0-14 14v28a6 6 0 0 0 6 6Zm136 92a6 6 0 0 1-4.8-2.4a54 54 0 0 0-86.4 0a6 6 0 1 1-9.6-7.2a65.65 65.65 0 0 1 29.69-22.26a38 38 0 1 1 46.22 0a65.65 65.65 0 0 1 29.69 22.26a6 6 0 0 1-4.8 9.6Zm-48-36a26 26 0 1 0-26-26a26 26 0 0 0 26 26Z"/>`),
		g.Group(children),
	)
}