package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Markdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M29.693 25.849H2.308a2.311 2.311 0 0 1-2.307-2.307V8.459a2.311 2.311 0 0 1 2.307-2.307h27.385A2.311 2.311 0 0 1 32 8.459v15.078a2.305 2.305 0 0 1-2.307 2.307zm-22-4.62v-6l3.078 3.849l3.073-3.849v6h3.078V10.771h-3.078l-3.073 3.849l-3.078-3.849H4.615v10.464zM28.307 16h-3.078v-5.229h-3.073V16h-3.078l4.615 5.385z"/>`),
		g.Group(children),
	)
}