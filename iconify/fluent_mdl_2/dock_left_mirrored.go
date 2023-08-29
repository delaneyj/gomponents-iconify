package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DockLeftMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384H0v1152h2048V384zm-128 128v896h-384V512h384zM128 1408V512h1280v896H128z"/>`),
		g.Group(children),
	)
}