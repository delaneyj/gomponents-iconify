package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 29a.999.999 0 0 1-.768-.36l-10-12a1 1 0 0 1 0-1.28l10-12A1 1 0 0 1 20 4v7.032c7.007.463 11 5.86 11 14.968a1 1 0 0 1-1.8.6c-2.822-3.762-5.391-5.346-9.2-5.571V28a1 1 0 0 1-1 1Zm-8.698-13L18 25.238V20a1 1 0 0 1 1-1a12.713 12.713 0 0 1 9.841 4.09C28.086 16.556 24.66 13 19 13a1 1 0 0 1-1-1V6.762Z"/><path fill="currentColor" d="M11.464 28.918L1.232 16.64a1 1 0 0 1 0-1.28L11.464 3.08L13 4.36L3.302 16L13 27.638Z"/>`),
		g.Group(children),
	)
}