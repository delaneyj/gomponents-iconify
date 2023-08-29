package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 21H4a1 1 0 0 1-1-1v-7.3a1 1 0 0 1 .341-.752L5 10.5V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1ZM9 7.329a1 1 0 0 1 .658.248l5 4.375A1 1 0 0 1 15 12.7V19h4V5H7v3.75l1.341-1.174A1 1 0 0 1 9 7.329ZM8 16h2v3h3v-5.843l-4-3.5l-4 3.5V19h3v-3Z"/>`),
		g.Group(children),
	)
}