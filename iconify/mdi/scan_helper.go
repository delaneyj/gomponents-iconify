package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanHelper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 2a2 2 0 0 1 2-2h4v2H2v4H0V2m24 20a2 2 0 0 1-2 2h-4v-2h4v-4h2v4M2 24a2 2 0 0 1-2-2v-4h2v4h4v2H2M22 0a2 2 0 0 1 2 2v4h-2V2h-4V0h4Z"/>`),
		g.Group(children),
	)
}