package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trees(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.74 16.319A4.995 4.995 0 0 1 10 17.9V21a1 1 0 1 1-2 0v-3.1A5.002 5.002 0 0 1 4 13V7a5 5 0 0 1 9.98-.453A4 4 0 0 1 20 10v4a4.002 4.002 0 0 1-3 3.874V21a1 1 0 1 1-2 0v-3.126a4.005 4.005 0 0 1-2.26-1.555ZM12 7v6a3.001 3.001 0 0 1-2 2.83V13a1 1 0 1 0-2 0v2.83A3.001 3.001 0 0 1 6 13V7a3 3 0 0 1 6 0Zm5 8.732V13a1 1 0 1 0-2 0v2.732A2 2 0 0 1 14 14v-4a2 2 0 1 1 4 0v4a2 2 0 0 1-1 1.732Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}