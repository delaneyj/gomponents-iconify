package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineOfflineShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 5H4v16c0 1.1.9 2 2 2h10v-2H6V5z"/><path fill="currentColor" d="M18 1h-8c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2zm0 16h-8v-1h8v1zm0-3h-8V6h8v8zm0-10h-8V3h8v1z"/><path fill="currentColor" d="M12.5 10.25h1.63l-.69.69L14.5 12L17 9.5L14.5 7l-1.06 1.06l.69.69H12c-.55 0-1 .45-1 1V12h1.5v-1.75z"/>`),
		g.Group(children),
	)
}