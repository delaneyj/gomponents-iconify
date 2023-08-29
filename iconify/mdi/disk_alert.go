package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiskAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M10 13.998a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm0-10a8 8 0 1 0 0 16a8 8 0 0 0 0-16zm10 8h2v-5h-2m0 9h2v-2h-2v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}