package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextRotateHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1664H0V128zm1920 1536V256H128v1408h1792zM436 896l-48 128H285l240-640h102l240 640H764l-48-128H436zm140-375l-92 247h184l-92-247zm803 922l163-163H256v-128h1286l-163-163l90-90l317 317l-317 317l-90-90z"/>`),
		g.Group(children),
	)
}