package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoundNoSim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.09 4.44a.996.996 0 0 0 0 1.41l2.03 2.03l-.12.13V19c0 1.1.9 2 2 2h10c.35 0 .68-.1.97-.26l1.17 1.17a.996.996 0 1 0 1.41-1.41L4.5 4.44a.996.996 0 0 0-1.41 0zM19 16.11V5c0-1.1-.9-2-2-2h-6.99L7.95 5.06L19 16.11z"/>`),
		g.Group(children),
	)
}