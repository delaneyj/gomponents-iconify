package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Addthis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 17.995h-6v5.979h-4v-5.979H8v-3.984h6V8.027h4v5.984h6zM28 .052H4a4.008 4.008 0 0 0-4 3.99v23.922a4.007 4.007 0 0 0 4 3.984h24a4.007 4.007 0 0 0 4-3.984V4.042a4.01 4.01 0 0 0-4-3.99z"/>`),
		g.Group(children),
	)
}