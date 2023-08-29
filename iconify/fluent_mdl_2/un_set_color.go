package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnSetColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1536H0V256h2048zm-128 128H128v1280h1792V384zm-128 1152H256V512h1536v1024zm-128-896H540l1124 715V640zM384 1408h1124L384 693v715z"/>`),
		g.Group(children),
	)
}