package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 128v896h896v1024H0V128h1024zM896 1920v-768H128v768h768zm0-896V256H128v768h768zm896 128h-768v768h768v-768zm-128-768h384v128h-384v384h-128V512h-384V384h384V0h128v384z"/>`),
		g.Group(children),
	)
}