package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Jamendo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M24 2a22 22 0 1 0 22 22A22 22 0 0 0 24 2Zm0 35.91A13.72 13.72 0 0 1 10.47 24A13.72 13.72 0 0 1 24 10.09A13.72 13.72 0 0 1 37.53 24A13.72 13.72 0 0 1 24 37.91Z"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="M31.63 24.9L21.05 32a1 1 0 0 1-1.4-.2a1 1 0 0 1-.2-.65V16.82a1 1 0 0 1 1-1a1 1 0 0 1 .64.21l10.58 7.2a1 1 0 0 1 .33 1.35a1.12 1.12 0 0 1-.37.32Z"/>`),
		g.Group(children),
	)
}