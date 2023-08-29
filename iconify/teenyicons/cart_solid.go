package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.979.356L.02.644L3.128 11H15V4.5A2.5 2.5 0 0 0 12.5 2H1.472L.979.356Z"/><path fill="currentColor" fill-rule="evenodd" d="M5.5 12a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3ZM5 13.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Zm7.5-1.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Zm-.5 1.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}