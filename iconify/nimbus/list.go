package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3.8 3.05H16v1.26H3.8zm0 4.29H16V8.6H3.8zm0 4.36H16v1.26H3.8zM1.25 2.5A1.22 1.22 0 0 1 2.5 3.68a1.22 1.22 0 0 1-1.25 1.17A1.22 1.22 0 0 1 0 3.68A1.22 1.22 0 0 1 1.25 2.5zm0 4.3a1.18 1.18 0 1 1 0 2.35a1.18 1.18 0 1 1 0-2.35zm0 4.35a1.18 1.18 0 1 1 0 2.35a1.18 1.18 0 1 1 0-2.35z"/>`),
		g.Group(children),
	)
}