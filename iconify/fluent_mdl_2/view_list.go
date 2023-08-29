package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v2048H0V0h2048zM128 640h512V128H128v512zm0 640h512V768H128v512zm512 128H128v512h512v-512zm1280 0H768v512h1152v-512zm0-640H768v512h1152V768zm0-128V128H768v512h1152z"/>`),
		g.Group(children),
	)
}