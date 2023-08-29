package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTypingAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0V15H7.5A7.5 7.5 0 0 1 0 7.5ZM4 8h1V7H4v1Zm7 0h-1V7h1v1ZM7 8h1V7H7v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}