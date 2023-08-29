package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosePaneMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384H0v1152h2048V384zm-128 128v896H640V512h1280zM128 1408V512h384v896H128zm931-765L742 960l317 317l90-90l-162-163h421V896H987l162-163l-90-90z"/>`),
		g.Group(children),
	)
}