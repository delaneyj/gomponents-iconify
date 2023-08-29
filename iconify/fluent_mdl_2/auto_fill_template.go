package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoFillTemplate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v640H0V256h2048zm-128 128H128v384h1792V384zM0 1024h2048v640H0v-640zm128 512h1792v-384H128v384zm1408-896H256V512h1280v128zM256 1280h640v128H256v-128z"/>`),
		g.Group(children),
	)
}