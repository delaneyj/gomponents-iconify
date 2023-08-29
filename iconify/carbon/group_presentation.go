package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupPresentation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M10 30H8v-3H4v3H2v-3a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2zm20 0h-2v-3h-4v3h-2v-3a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2zm-10 0h-2v-3h-4v3h-2v-3a2.002 2.002 0 0 1 2-2h4a2.002 2.002 0 0 1 2 2z"/><circle cx="16" cy="22" r="2" fill="currentColor"/><circle cx="6" cy="22" r="2" fill="currentColor"/><circle cx="26" cy="22" r="2" fill="currentColor"/><circle cx="21" cy="18" r="2" fill="currentColor"/><circle cx="11" cy="18" r="2" fill="currentColor"/><path fill="currentColor" d="M26 14H6a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h20a2.002 2.002 0 0 1 2 2v8a2.002 2.002 0 0 1-2 2ZM6 4v8h20V4Z"/>`),
		g.Group(children),
	)
}