package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberedListTextMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1536v-128h1536v128H0zm0-512V896h1536v128H0zm0-640h1536v128H0V384z"/>`),
		g.Group(children),
	)
}