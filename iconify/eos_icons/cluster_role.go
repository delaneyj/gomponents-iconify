package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClusterRole(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 4h-4.18a3 3 0 0 0-5.64 0H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Zm-7 0a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm7 15H5v-4h14Zm0-6H5V9h14Z"/><circle cx="17" cy="11" r="1" fill="currentColor"/><circle cx="14" cy="11" r="1" fill="currentColor"/><circle cx="14" cy="17" r="1" fill="currentColor"/><circle cx="17" cy="17" r="1" fill="currentColor"/><path fill="currentColor" d="M6 10h5v2H6zm0 6h5v2H6z"/>`),
		g.Group(children),
	)
}