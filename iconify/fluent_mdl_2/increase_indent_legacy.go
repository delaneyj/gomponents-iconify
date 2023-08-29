package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IncreaseIndentLegacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 384h1024v128H1024V384zm0 1152v-128h1024v128H1024zm0-512V896h1024v128H1024zm768-384v128h-768V640h768zm0 512v128h-768v-128h768zM483 733l90-90l317 317l-317 317l-90-90l163-163H0V896h646L483 733z"/>`),
		g.Group(children),
	)
}