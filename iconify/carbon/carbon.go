package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Carbon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13.5 30.815a1.001 1.001 0 0 1-.493-.13l-8.5-4.815A1 1 0 0 1 4 25V15a1 1 0 0 1 .507-.87l8.5-4.815a1.001 1.001 0 0 1 .986 0l8.5 4.815A1 1 0 0 1 23 15v10a1 1 0 0 1-.507.87l-8.5 4.815a1.001 1.001 0 0 1-.493.13ZM6 24.417l7.5 4.248l7.5-4.248v-8.834l-7.5-4.248L6 15.582Z"/><path fill="currentColor" d="M28 17h-2V7.583l-7.5-4.248l-8.007 4.535l-.986-1.74l8.5-4.815a1.001 1.001 0 0 1 .986 0l8.5 4.815A1 1 0 0 1 28 7Z"/>`),
		g.Group(children),
	)
}