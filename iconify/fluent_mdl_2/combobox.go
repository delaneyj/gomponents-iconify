package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Combobox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1408H0V256h2048zm-128 128H128v1152h1792V384zm-384 768l-256-256h512l-256 256zM512 1024H256V896h256v128zm512 0H768V896h256v128z"/>`),
		g.Group(children),
	)
}