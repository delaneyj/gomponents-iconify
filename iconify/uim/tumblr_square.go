package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TumblrSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.571 17.85a3.396 3.396 0 0 1-3.747-3.332v-4.076H8.562v-1.61a3.817 3.817 0 0 0 2.454-3.363a.099.099 0 0 1 .09-.093h1.827v3.178h2.496v1.888h-2.505v3.884a1.094 1.094 0 0 0 1.2 1.229a2.936 2.936 0 0 0 1.007-.215l.6 1.779c-.611.49-1.376.75-2.16.73z" opacity=".5"/><path fill="currentColor" d="M2.019 2v20h19.963V2H2.019zm11.552 15.85a3.396 3.396 0 0 1-3.747-3.332v-4.076H8.562v-1.61a3.817 3.817 0 0 0 2.454-3.363a.099.099 0 0 1 .09-.093h1.827v3.178h2.496v1.888h-2.505v3.884a1.094 1.094 0 0 0 1.2 1.229a2.936 2.936 0 0 0 1.007-.215l.6 1.779c-.611.49-1.376.75-2.16.73z"/>`),
		g.Group(children),
	)
}