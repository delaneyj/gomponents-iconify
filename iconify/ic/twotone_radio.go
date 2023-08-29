package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneRadio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 13H4v7h16v-7zM8 18.98a2.5 2.5 0 0 1 0-5a2.5 2.5 0 0 1 0 5z" opacity=".3"/><path fill="currentColor" d="M2 20a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8c0-1.11-.89-2-2-2H8.3l8.26-3.34L15.88 1L3.24 6.15C2.51 6.43 2 7.17 2 8v12zM4 8h16v3h-2V9h-2v2H4V8zm0 5h16v7H4v-7z"/><circle cx="8" cy="16.48" r="2.5" fill="currentColor"/>`),
		g.Group(children),
	)
}