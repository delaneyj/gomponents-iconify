package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenPaneMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384H0v1152h2048V384zm-128 128v896H640V512h1280zM128 1408V512h384v896H128zm640-512v128h421l-162 163l90 90l317-317l-317-317l-90 90l162 163H768z"/>`),
		g.Group(children),
	)
}