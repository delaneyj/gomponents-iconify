package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneVideoChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17.17L5.17 16H20V4H4v13.17zM7 7c0-.55.45-1 1-1h6c.55 0 1 .45 1 1v1.99L17 7v6l-2-1.99V13c0 .55-.45 1-1 1H8c-.55 0-1-.45-1-1V7z" opacity=".3"/><path fill="currentColor" d="M20 2H4c-1.1 0-1.99.9-1.99 2L2 22l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm0 14H5.17L4 17.17V4h16v12z"/><path fill="currentColor" d="M8 14h6c.55 0 1-.45 1-1v-1.99L17 13V7l-2 1.99V7c0-.55-.45-1-1-1H8c-.55 0-1 .45-1 1v6c0 .55.45 1 1 1z"/>`),
		g.Group(children),
	)
}