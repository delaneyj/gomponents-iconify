package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DecreaseIndentLegacy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 384h1024v128H1024V384zm0 1152v-128h1024v128H1024zm0-512V896h1024v128H1024zm768-384v128h-768V640h768zm0 512v128h-768v-128h768zM413 733L250 896h646v128H250l163 163l-90 90L6 960l317-317l90 90z"/>`),
		g.Group(children),
	)
}