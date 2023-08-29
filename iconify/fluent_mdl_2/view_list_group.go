package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewListGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v2048H384V768H0V0h2048zM640 640V128H128v512h512zm-128 640h512V768H512v512zm512 128H512v512h512v-512zm896 0h-768v512h768v-512zm0-640h-768v512h768V768zm0-128V128H768v512h1152z"/>`),
		g.Group(children),
	)
}