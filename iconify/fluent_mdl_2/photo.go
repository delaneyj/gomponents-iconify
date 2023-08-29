package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Photo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 256h2048v1536H0V256zm1920 1408V384H128v1280h1792zM1792 512v1024H256V512h1536zm-128 896V640H384v768h1280z"/>`),
		g.Group(children),
	)
}