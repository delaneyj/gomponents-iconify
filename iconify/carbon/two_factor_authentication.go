package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoFactorAuthentication(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m11 23.18l-2-2.001l-1.411 1.41L11 26l6-6l-1.41-1.41L11 23.18zM28 30h-4v-2h4V16h-4V8a4.005 4.005 0 0 0-4-4V2a6.007 6.007 0 0 1 6 6v6h2a2.002 2.002 0 0 1 2 2v12a2.002 2.002 0 0 1-2 2z"/><path fill="currentColor" d="M20 14h-2V8A6 6 0 0 0 6 8v6H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V16a2 2 0 0 0-2-2ZM8 8a4 4 0 0 1 8 0v6H8Zm12 20H4V16h16Z"/>`),
		g.Group(children),
	)
}