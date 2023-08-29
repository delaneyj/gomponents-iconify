package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4c-4.879 0-9 4.121-9 9s4.121 9 9 9s9-4.121 9-9s-4.121-9-9-9zm4 10h-3v3h-2v-3H8v-2h3V9h2v3h3v2zm1.284-10.293l1.412-1.416l3.01 3l-1.413 1.417zM5.282 2.294L6.7 3.706l-2.99 3l-1.417-1.413z"/>`),
		g.Group(children),
	)
}