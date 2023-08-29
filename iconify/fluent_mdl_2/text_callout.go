package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextCallout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1408H0V256h2048zm-128 128H128v1152h1792V384zm-256 640h-640V896h640v128zm-896 0H384V896h384v128z"/>`),
		g.Group(children),
	)
}