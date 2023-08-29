package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.84 22.73L18.11 20H5v-2h11.11l-3.05-3.05L12 16L5 9h2.11l-6-6l1.28-1.27l19.72 19.73l-1.27 1.27M19 9h-4V3H9v2.8l6.6 6.6L19 9Z"/>`),
		g.Group(children),
	)
}