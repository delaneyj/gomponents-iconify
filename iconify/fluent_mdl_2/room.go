package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Room(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 128v1664h-442l-473 95l-237 47V779l640-128v-11H640v1152H128V128h1792zM896 885v893l512-103V782L896 885zm896 779V256H256v1408h256V512h1024v1152h256z"/>`),
		g.Group(children),
	)
}