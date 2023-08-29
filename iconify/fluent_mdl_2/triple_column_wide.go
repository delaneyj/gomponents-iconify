package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TripleColumnWide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 256h512v1536H0V256zm128 1408h256V384H128v1280zM2048 256v1536h-512V256h512zm-128 128h-256v1280h256V384zM640 256h768v1536H640V256zm128 1408h512V384H768v1280z"/>`),
		g.Group(children),
	)
}