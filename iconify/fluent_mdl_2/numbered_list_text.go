package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberedListText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M512 1536v-128h1536v128H512zm0-1152h1536v128H512V384zm0 640V896h1536v128H512z"/>`),
		g.Group(children),
	)
}