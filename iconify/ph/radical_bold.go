package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadicalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M244 72v24a12 12 0 0 1-24 0V84h-91.68L83.24 204.21a12 12 0 0 1-22.47 0l-48-128a12 12 0 1 1 22.47-8.43l36.76 98l36.77-98A12 12 0 0 1 120 60h112a12 12 0 0 1 12 12Z"/>`),
		g.Group(children),
	)
}