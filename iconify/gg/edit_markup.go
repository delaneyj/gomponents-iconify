package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditMarkup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 24c6.627 0 12-5.373 12-12S18.627 0 12 0S0 5.373 0 12s5.373 12 12 12Zm6.58-4.469A9.976 9.976 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12a9.975 9.975 0 0 0 3.333 7.453L7 10.973h2.17l2.83-4.9l2.83 4.9H17l1.58 8.558Zm-2.488 1.596l-.886-8.153H8.794l-.886 8.153A9.964 9.964 0 0 0 12 22a9.965 9.965 0 0 0 4.092-.873Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}