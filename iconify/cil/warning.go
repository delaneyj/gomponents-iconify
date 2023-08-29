package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M240 176h32v176h-32zm0 208h32v32h-32z"/><path fill="currentColor" d="M274.014 16h-36.028L16 445.174V496h480v-50.826ZM464 464H48v-11.041L256 50.826l208 402.133Z"/>`),
		g.Group(children),
	)
}