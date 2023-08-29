package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualRealityOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.826 10a3.012 3.012 0 0 0-1.955-1.871A3.005 3.005 0 0 0 16 6H8a3.005 3.005 0 0 0-2.871 2.129A3.011 3.011 0 0 0 3.174 10H2v4h1.174a3.011 3.011 0 0 0 1.955 1.871A3.006 3.006 0 0 0 8 18h1.605l1.891-3.275a.582.582 0 0 1 1.008 0L14.396 18H16a3.006 3.006 0 0 0 2.871-2.129A3.012 3.012 0 0 0 20.826 14H22v-4ZM19 13a1 1 0 0 1-1 1h-1v1a1 1 0 0 1-1 1h-.45l-1.314-2.275a2.582 2.582 0 0 0-4.472 0L8.45 16H8a1.001 1.001 0 0 1-1-1v-1H6a1.001 1.001 0 0 1-1-1v-2a1.001 1.001 0 0 1 1-1h1V9a1.001 1.001 0 0 1 1-1h8a1.001 1.001 0 0 1 1 1v1h1a1.001 1.001 0 0 1 1 1Z"/><path fill="currentColor" d="M15 11H9v-1h6Z"/>`),
		g.Group(children),
	)
}