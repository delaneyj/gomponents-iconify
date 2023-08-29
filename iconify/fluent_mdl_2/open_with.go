package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenWith(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2045 1219l-317-317l-317 317l90 90l162-163v583h129v-583l163 163l90-90zm-253-834H639V256h1153v129zm-256 1280H639v-129h897v129zM384 512H0V128h384v384zM128 384h128V256H128v128zm256 768H0V768h384v384zm-256-128h128V896H128v128zm256 768H0v-384h384v384zm-256-128h128v-128H128v128zm1280-640H640V896h896l-128 128z"/>`),
		g.Group(children),
	)
}