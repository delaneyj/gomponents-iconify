package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolcanoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 8h-7l-2 5H6l-4 9h20L18 8M7.3 15h3.05l.5-1.26l1.5-3.74h4.15l2.85 10H5.08l2.22-5M13 1h2v4h-2V1m3.12 4.47l2.83-2.83l1.41 1.41l-2.82 2.83l-1.42-1.41M7.64 4.05l1.41-1.41l2.83 2.82l-1.41 1.42l-2.83-2.83Z"/>`),
		g.Group(children),
	)
}