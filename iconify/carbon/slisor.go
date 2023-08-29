package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slisor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 20h-2v2h2v6H4v-6h2v-2H4a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2Z"/><circle cx="7" cy="25" r="1" fill="currentColor"/><path fill="currentColor" d="M16 22a1 1 0 0 1-.447-.105l-10-5l.894-1.79L16 19.883l9.553-4.777l.894 1.79l-10 5A1 1 0 0 1 16 22Z"/><path fill="currentColor" d="M16 16a1 1 0 0 1-.447-.105l-10-5a1 1 0 0 1 0-1.79l10-5a1 1 0 0 1 .894 0l10 5a1 1 0 0 1 0 1.79l-10 5A1 1 0 0 1 16 16Zm-7.764-6L16 13.882L23.764 10L16 6.118Z"/>`),
		g.Group(children),
	)
}