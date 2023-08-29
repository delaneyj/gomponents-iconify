package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Master(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20h16v2H4v-2ZM21 5a1 1 0 0 0-1 1a.899.899 0 0 0 .064.332a1.104 1.104 0 0 0 .167.283l-2.12 1.692L15.994 10L14.24 6.927l-1.752-3.072a1.045 1.045 0 0 0 .369-.353A.946.946 0 0 0 13 3a1 1 0 1 0-2 0a.946.946 0 0 0 .143.502a1.046 1.046 0 0 0 .369.353L9.757 6.927L8.002 10L5.886 8.308L3.769 6.615a1.104 1.104 0 0 0 .167-.283A.899.899 0 0 0 4 6a1 1 0 1 0-1 1a.147.147 0 0 0 .041-.007a.212.212 0 0 1 .041-.01l.459 5.509L4 18h16l.459-5.508l.459-5.509a.212.212 0 0 1 .04.01A.147.147 0 0 0 21 7a1 1 0 1 0 0-2Z"/>`),
		g.Group(children),
	)
}