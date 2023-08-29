package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuadColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 256h384v1536H0V256zm128 1408h128V384H128v1280zm896-1408h384v1536h-384V256zm128 1408h128V384h-128v1280zM512 256h384v1536H512V256zm128 1408h128V384H640v1280zM1920 256v1536h-384V256h384zm-128 128h-128v1280h128V384z"/>`),
		g.Group(children),
	)
}